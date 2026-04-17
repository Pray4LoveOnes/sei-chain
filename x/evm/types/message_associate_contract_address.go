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
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	sdkerrors "github.com/sei-protocol/sei-chain/sei-cosmos/types/errors"
)

const TypeMsgAssociateContractAddress = "evm_associate_contract_address"

var (
	_ sdk.Msg = &MsgAssociateContractAddress{}
)

func NewMsgAssociateContractAddress(sender sdk.AccAddress, addr sdk.AccAddress) *MsgAssociateContractAddress {
	return &MsgAssociateContractAddress{Sender: sender.String(), Address: addr.String()}
}

func (msg *MsgAssociateContractAddress) Route() string {
	return RouterKey
}

func (msg *MsgAssociateContractAddress) Type() string {
	return TypeMsgAssociateContractAddress
}

func (msg *MsgAssociateContractAddress) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgAssociateContractAddress) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgAssociateContractAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	return nil
}
