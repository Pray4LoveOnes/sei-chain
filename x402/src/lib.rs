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

//! # x402 SovereignKin Orchestrator
//! 
//! This crate provides the high-performance orchestration layer for the SovereignKin network,
//! focusing on lock-stripped payment batching and native pre-validation before Wasm execution.

pub mod payment_batcher;

// Re-export the primary components for a cleaner public API
pub use payment_batcher::{Payment, PaymentBatcher};
