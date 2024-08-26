package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateAgentAdmin{}

// ValidateBasic performs stateless validation on MsgUpdateAgentAdmin.
func (m MsgUpdateAgentAdmin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.NewAdmin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new admin: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateAgentAdmin.
func (m MsgUpdateAgentAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateAgentAdmin) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateAgentAdmin) Type() string {
	return sdk.MsgTypeURL(&m)
}