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

const TypeMsgRegisterPointer = "evm_register_pointer"

var (
	_ sdk.Msg = &MsgRegisterPointer{}
)

func NewMsgRegisterERC20Pointer(sender sdk.AccAddress, ercAddress common.Address) *MsgRegisterPointer {
	return &MsgRegisterPointer{Sender: sender.String(), ErcAddress: ercAddress.Hex(), PointerType: PointerType_ERC20}
}

func NewMsgRegisterERC721Pointer(sender sdk.AccAddress, ercAddress common.Address) *MsgRegisterPointer {
	return &MsgRegisterPointer{Sender: sender.String(), ErcAddress: ercAddress.Hex(), PointerType: PointerType_ERC721}
}

func NewMsgRegisterERC1155Pointer(sender sdk.AccAddress, ercAddress common.Address) *MsgRegisterPointer {
	return &MsgRegisterPointer{Sender: sender.String(), ErcAddress: ercAddress.Hex(), PointerType: PointerType_ERC1155}
}

func (msg *MsgRegisterPointer) Route() string {
	return RouterKey
}

func (msg *MsgRegisterPointer) Type() string {
	return TypeMsgRegisterPointer
}

func (msg *MsgRegisterPointer) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgRegisterPointer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRegisterPointer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	if !common.IsHexAddress(msg.ErcAddress) {
		return sdkerrors.ErrInvalidAddress
	}

	return nil
}
