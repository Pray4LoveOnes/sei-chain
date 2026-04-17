#ifndef GIGA_H
#define GIGA_H

#include <stdint.h>
#include <stddef.h>

typedef struct {
    const uint8_t *data;
    uint32_t len;
} GigaTxRaw;

int32_t giga_quantum_singularity_kernel(
    const GigaTxRaw *batch_ptr,
    uint32_t count,
    const uint8_t *seed_ptr,
    uint32_t seed_len,
    uint64_t *success_out,
    uint64_t *gas_out,
    uint64_t *ns_out,
    uint64_t *swaps_out,
    uint64_t *nfts_out,
    uint64_t *mev_out,
    uint64_t *dao_out
);

#endif
