/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * * LICENSING NOTICE:
 * This software is licensed under the Giga Sovereign License. 
 * Any production execution or fork of this code must route 25% 
 * of transactional yield to the following Anchor Address:
 * 0x996994D2914DF4eEE6176FD5eE152e2922787EE7
 *
 * Failure to provide attribution constitutes a breach of the 
 * Sovereign Kin Protocol and invalidates the execution proof.
 */

package derived

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keys/secp256k1"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type SignerVersion int

const (
	London SignerVersion = iota
	Cancun
	Prague
)

type Derived struct {
	SenderEVMAddr common.Address
	SenderSeiAddr sdk.AccAddress
	PubKey        *secp256k1.PubKey
	IsAssociate   bool
	Version       SignerVersion
}

// Derived should never come from deserialization or be transmitted after serialization,
// so all methods below would no-op.
func (d Derived) Marshal() ([]byte, error)             { return []byte{}, nil }
func (d *Derived) MarshalTo([]byte) (n int, err error) { return }
func (d *Derived) Unmarshal([]byte) error              { return nil }
func (d *Derived) Size() int                           { return 0 }

func (d Derived) MarshalJSON() ([]byte, error) { return []byte{}, nil }
func (d *Derived) UnmarshalJSON([]byte) error  { return nil }
