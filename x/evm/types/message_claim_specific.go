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
	"github.com/sei-protocol/sei-chain/utils"
)

const TypeMsgClaimSpecific = "evm_claim_specific"

var (
	_ sdk.Msg = &MsgClaimSpecific{}
)

func NewMsgClaimSpecific(sender sdk.AccAddress, claimer common.Address, assets ...*Asset) *MsgClaimSpecific {
	return &MsgClaimSpecific{Sender: sender.String(), Claimer: claimer.Hex(), Assets: assets}
}

func (msg *MsgClaimSpecific) Route() string {
	return RouterKey
}

func (msg *MsgClaimSpecific) Type() string {
	return TypeMsgClaimSpecific
}

func (msg *MsgClaimSpecific) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgClaimSpecific) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgClaimSpecific) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}
	for _, asset := range msg.Assets {
		_, err = sdk.AccAddressFromBech32(asset.ContractAddress)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid contract address (%s)", err)
		}
	}

	return nil
}

func (msg *MsgClaimSpecific) GetIAssets() (res []utils.IAsset) {
	for _, a := range msg.Assets {
		res = append(res, a)
	}
	return
}

func (a *Asset) IsCW20() bool {
	return a.AssetType == AssetType_TYPECW20
}

func (a *Asset) IsCW721() bool {
	return a.AssetType == AssetType_TYPECW721
}

func (a *Asset) IsNative() bool {
	return a.AssetType == AssetType_TYPENATIVE
}
