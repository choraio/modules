package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSubject{}

// ValidateBasic performs stateless validation on MsgCreateSubject.
func (m MsgCreateSubject) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Steward); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("steward: %s", err)
	}

	if m.Metadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("metadata: empty string is not allowed")
	}

	if len(m.Metadata) > MetadataMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateSubject.
func (m MsgCreateSubject) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Steward)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgCreateSubject) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgCreateSubject) Type() string {
	return sdk.MsgTypeURL(&m)
}