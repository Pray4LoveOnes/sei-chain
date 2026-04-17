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

// build.rs — UNTOUCHABLE SINGULARITY V2 (AVX2 + Rayon + Zero-Heap Hot Path)

//use std::env;
//use std::path::Path;

fn main() {
    println!("cargo:rerun-if-changed=src/lib.rs");
    println!("cargo:rerun-if-changed=build.rs");

    // Force native CPU features for maximum speed on your MBP
    println!("cargo:rustc-env=RUSTFLAGS=-C target-cpu=native -C target-feature=+avx2,+bmi2");

    // Optional: Link-time optimization for the final binary
    println!("cargo:rustc-link-arg=-flto");
}
