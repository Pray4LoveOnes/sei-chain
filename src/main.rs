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

use std::sync::Arc;
use std::time::Instant;
use rayon::prelude::*;
use colored::Colorize;
use ed25519_dalek::{SigningKey, Verifier, Signer, SECRET_KEY_LENGTH};
use rand::rngs::OsRng;
use rand::RngCore;
use giga_native::{GigaTxRaw, giga_quantum_singularity_kernel};

#[tokio::main]
async fn main() -> eyre::Result<()> {
    let mut success: u64 = 0;
    let mut gas: u64 = 0;
    let mut ns: u64 = 0;
    let mut swaps: u64 = 0;
    let mut nfts: u64 = 0;
    let mut mev: u64 = 0;
    let mut dao: u64 = 0;

    println!("{}", "IGNITING NATIVE SEI GIGA KERNEL...".bright_cyan().bold());

    let mut seed = [0u8; SECRET_KEY_LENGTH];
    OsRng.fill_bytes(&mut seed);
    let signing_key = SigningKey::from_bytes(&seed);
    let verifying_key = signing_key.verifying_key();
    let signing_key_arc = Arc::new(signing_key);

    let batch_count = 777_777;
    let start_compute = Instant::now();

    let signatures: Vec<Vec<u8>> = (0..batch_count)
        .into_par_iter()
        .map(|i| {
            let msg = format!("SEI-TX-{}", i);
            let sig = signing_key_arc.sign(msg.as_bytes());
            sig.to_bytes().to_vec()
        })
        .collect();

    let compute_elapsed = start_compute.elapsed();

    // Fix E0308: Cast len to u32
    let ffi_batch: Vec<GigaTxRaw> = signatures
        .iter()
        .map(|s| GigaTxRaw {
            data: s.as_ptr(),
            len: s.len() as u32,
        })
        .collect();

    let start_ffi = Instant::now();

    // Fix E0061: Provide all 11 arguments
    unsafe {
        giga_quantum_singularity_kernel(
            ffi_batch.as_ptr(),
            ffi_batch.len() as u32,
            seed.as_ptr(),
            seed.len() as u32,
            &mut success,
            &mut gas,
            &mut ns,
            &mut swaps,
            &mut nfts,
            &mut mev,
            &mut dao,
        );
    }

    let ffi_elapsed = start_ffi.elapsed();
    let total_elapsed = compute_elapsed + ffi_elapsed;
    let real_tps = batch_count as f64 / total_elapsed.as_secs_f64();

    println!("\n{}", "--- NATIVE PERFORMANCE REPORT ---".bright_green());
    println!("{:<25} {:?}", "FFI Kernel Time:", ffi_elapsed);
    println!("{:<25} {} TPS", "Real Throughput:", format!("{:.2}", real_tps).bright_white());
    println!("{:<25} {}", "Total TX Handled:", success.to_string().bright_yellow());

    Ok(())
}
