package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateMetadata{}

// ValidateBasic performs stateless validation on MsgUpdateMetadata.
func (m MsgUpdateMetadata) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	if m.NewMetadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("new metadata: empty string is not allowed")
	}

	if len(m.NewMetadata) > MetadataMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("new metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateMetadata.
func (m MsgUpdateMetadata) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Address)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateMetadata) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateMetadata) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateMetadata) Type() string {
	return sdk.MsgTypeURL(&m)
}