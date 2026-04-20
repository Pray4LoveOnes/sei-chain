package evmrpc

import (
	"runtime"
	"sync/atomic"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// [SOVEREIGN-GIGA] MEGA-STACK KERNEL || VERSION: ζΩ∞ v4.1 (SHARDED)
// ---------------------------------------------------------------------------------------------------------------------

const (
	SlotRequests      = 0
	SlotSuccess       = 1
	SlotBlocksFetched = 2
	SlotTotalLatency  = 3
	SlotWSConnections = 4

	IdxFilterLogMatch    = 5
	IdxFilterBlockMatch  = 6
	IdxFilterPollLatency = 7

	MaxMetrics = 8
	CacheLine  = 64
)

var telemetryEnabled uint32 = 1

func TelemetryEnabled() bool {
	return atomic.LoadUint32(&telemetryEnabled) == 1
}

type MetricCell struct {
	Value uint64
	_pad  [CacheLine - 8]byte
}

type SovereignRegistry struct {
	_     [CacheLine]byte // Front-padding to isolate from slice header
	cells [MaxMetrics]MetricCell
}

type ShardedRegistry struct {
	shards  []SovereignRegistry
	nextIdx uint64
}

var GlobalRegistry = NewShardedRegistry()

func NewShardedRegistry() *ShardedRegistry {
	n := runtime.GOMAXPROCS(0)
	if n < 1 {
		n = 1
	}
	return &ShardedRegistry{shards: make([]SovereignRegistry, n)}
}

func (s *ShardedRegistry) shard() *SovereignRegistry {
	n := uint64(len(s.shards))
	idx := atomic.AddUint64(&s.nextIdx, 1)
	return &s.shards[idx%n]
}

func (s *ShardedRegistry) RecordRequest(start time.Time, success bool) {
	if !TelemetryEnabled() {
		return
	}
	r := s.shard()
	atomic.AddUint64(&r.cells[SlotRequests].Value, 1)
	if success {
		atomic.AddUint64(&r.cells[SlotSuccess].Value, 1)
	}
	if !start.IsZero() {
		atomic.AddUint64(&r.cells[SlotTotalLatency].Value, uint64(time.Since(start)))
	}
}

func (s *ShardedRegistry) RecordEvent(slot int) {
	if !TelemetryEnabled() || slot < 0 || slot >= MaxMetrics {
		return
	}
	r := s.shard()
	atomic.AddUint64(&r.cells[slot].Value, 1)
}

func (s *ShardedRegistry) aggregate(slot int) uint64 {
	if slot < 0 || slot >= MaxMetrics {
		return 0
	}
	var total uint64
	for i := range s.shards {
		total += atomic.LoadUint64(&s.shards[i].cells[slot].Value)
	}
	return total
}

func (s *ShardedRegistry) Export(slot int) uint64 {
	return s.aggregate(slot)
}
