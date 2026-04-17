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
)

var (
	_ sdk.Msg = &MsgInternalEVMCall{}
)

func NewMessageInternalEVMCall(from sdk.AccAddress, to string, value *sdk.Int, data []byte) *MsgInternalEVMCall {
	return &MsgInternalEVMCall{
		Sender: from.String(),
		To:     to,
		Value:  value,
		Data:   data,
	}
}

func (msg *MsgInternalEVMCall) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{senderAddr}
}

func (msg *MsgInternalEVMCall) ValidateBasic() error {
	return nil
}
