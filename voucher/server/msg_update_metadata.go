package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/voucher/types/v1"
)

// UpdateMetadata implements the Msg/UpdateMetadata method.
func (s Server) UpdateMetadata(ctx context.Context, req *v1.MsgUpdateMetadata) (*v1.MsgUpdateMetadataResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from issuer address
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err // internal error
	}

	// get voucher from voucher table
	voucher, err := s.ss.VoucherTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("voucher with id %d: %s", req.Id, err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	voucherIssuer := sdk.AccAddress(voucher.Issuer)

	// verify issuer is voucher issuer
	if !voucherIssuer.Equals(issuer) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"issuer %s: voucher issuer %s", issuer, voucherIssuer.String(),
		)
	}

	// set new voucher metadata
	voucher.Metadata = req.NewMetadata

	// update voucher in voucher table
	err = s.ss.VoucherTable().Update(ctx, voucher)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateMetadata{
		Id: voucher.Id,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateMetadataResponse{
		Id: voucher.Id,
	}, nil
}
