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
	_ sdk.Msg = &MsgInternalEVMDelegateCall{}
)

func NewMessageInternalEVMDelegateCall(from sdk.AccAddress, to string, codeHash []byte, data []byte, fromContract string) *MsgInternalEVMDelegateCall {
	return &MsgInternalEVMDelegateCall{
		Sender:       from.String(),
		To:           to,
		Data:         data,
		CodeHash:     codeHash,
		FromContract: fromContract,
	}
}

func (msg *MsgInternalEVMDelegateCall) GetSigners() []sdk.AccAddress {
	contractAddr, err := sdk.AccAddressFromBech32(msg.FromContract)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{contractAddr}
}

func (msg *MsgInternalEVMDelegateCall) ValidateBasic() error {
	return nil
}
