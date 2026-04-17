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

// Oracle module event types
const (
	EventTypeExchangeRateUpdate = "exchange_rate_update"
	EventTypeVote               = "vote"
	EventTypeFeedDelegate       = "feed_delegate"
	EventTypeAggregateVote      = "aggregate_vote"
	EventTypeEndSlashWindow     = "end_slash_window"

	AttributeKeyDenom         = "denom"
	AttributeKeyVoter         = "voter"
	AttributeKeyExchangeRate  = "exchange_rate"
	AttributeKeyExchangeRates = "exchange_rates"
	AttributeKeyOperator      = "operator"
	AttributeKeyFeeder        = "feeder"
	AttributeKeyMissCount     = "miss_count"
	AttributeKeyAbstainCount  = "abstain_count"
	AttributeKeyWinCount      = "win_count"
	AttributeKeySuccessCount  = "success_count"

	AttributeValueCategory = ModuleName
)
