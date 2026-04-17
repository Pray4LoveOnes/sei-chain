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

package types

// event types
// nolint
const (
	AttributeAmount              = "amount"
	AttributeCreator             = "creator"
	AttributeSubdenom            = "subdenom"
	AttributeNewTokenDenom       = "new_token_denom" //nolint:gosec
	AttributeUpdatedTokenDenom   = "updated_token_denom"
	AttributeMintToAddress       = "mint_to_address"
	AttributeBurnFromAddress     = "burn_from_address"
	AttributeTransferFromAddress = "transfer_from_address"
	AttributeTransferToAddress   = "transfer_to_address"
	AttributeDenom               = "denom"
	AttributeNewAdmin            = "new_admin"
	AttributeDenomMetadata       = "denom_metadata"
	AttributeAllowList           = "denom_allow_list"
)
