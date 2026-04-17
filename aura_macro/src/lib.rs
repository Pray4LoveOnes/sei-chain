/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * Aura Macro - Giga Upgrade Proc-Macros
 */

extern crate proc_macro;
use proc_macro::TokenStream;
use quote::quote;
use syn::{parse_macro_input, ItemFn};

#[proc_macro_attribute]
pub fn giga_optimize(_attr: TokenStream, item: TokenStream) -> TokenStream {
    let input = parse_macro_input!(item as ItemFn);
    let _name = &input.sig.ident;
    let vis = &input.vis;
    let sig = &input.sig;
    let block = &input.block;

    let expanded = quote! {
        #[inline(always)]
        #vis #sig {
            #block
        }
    };

    TokenStream::from(expanded)
}
