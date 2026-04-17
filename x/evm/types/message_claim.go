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

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	sdkerrors "github.com/sei-protocol/sei-chain/sei-cosmos/types/errors"
)

const TypeMsgClaim = "evm_claim"

var (
	_ sdk.Msg = &MsgClaim{}
)

func NewMsgClaim(sender sdk.AccAddress, claimer common.Address) *MsgClaim {
	return &MsgClaim{Sender: sender.String(), Claimer: claimer.Hex()}
}

func (msg *MsgClaim) Route() string {
	return RouterKey
}

func (msg *MsgClaim) Type() string {
	return TypeMsgClaim
}

func (msg *MsgClaim) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	return nil
}
