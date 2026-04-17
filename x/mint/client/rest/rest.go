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

package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	govrest "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client/rest"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
	"github.com/sei-protocol/sei-chain/x/mint/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	clientrest "github.com/sei-protocol/sei-chain/sei-cosmos/client/rest"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client/tx"
	typesrest "github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := clientrest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
}

// PlanRequest defines a proposal for a new upgrade plan.
type UpdateMinterRequest struct {
	BaseReq     typesrest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string            `json:"title" yaml:"title"`
	Description string            `json:"description" yaml:"description"`
	Deposit     sdk.Coins         `json:"deposit" yaml:"deposit"`
	Minter      types.Minter      `json:"minter" yaml:"minter"`
}

func UpdateResourceDependencyProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "update_minter",
		Handler:  newUpdateMinterPostHandler(clientCtx),
	}
}

func newUpdateMinterPostHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateMinterRequest

		if !typesrest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if typesrest.CheckBadRequestError(w, err) {
			return
		}

		content := types.NewUpdateMinterProposalHandler(
			req.Title, req.Description, req.Minter,
		)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if typesrest.CheckBadRequestError(w, err) {
			return
		}
		if typesrest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
