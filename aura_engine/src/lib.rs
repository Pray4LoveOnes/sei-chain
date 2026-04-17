/*
 * Copyright (c) 2026 The Keeper
 * Aura Engine - Sovereign Kin Execution Layer
 */

use giga_x402::payment_batcher::{Payment, PaymentBatcher};

pub struct AuraEngine {
    pub batcher: PaymentBatcher,
}

impl AuraEngine {
    pub fn new(capacity: usize) -> Self {
        Self {
            batcher: PaymentBatcher::new(capacity),
        }
    }

    pub fn process_chain_event(&self, id: u64, amount: u128, recipient: String) {
        let payment = Payment {
            id,
            amount,
            recipient,
            signature: vec![], 
        };
        
        if let Err(e) = self.batcher.add_payment(payment) {
            eprintln!("Aura Engine Pressure Warning: {}", e);
        }
    }
}
