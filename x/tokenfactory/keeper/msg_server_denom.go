package keeper

import (
	"context"

	"monetachain/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func findString(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Denom already exists")
	}

	params := k.GetParams(ctx)
	if !params.CanAnyoneAddCoins && !findString(params.AccessList, msg.Owner) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Owner is not in the access list")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             msg.Ticker,
		Precision:          msg.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		Supply:             math.ZeroInt(),
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Denom to update not found")
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if !valFound.CanChangeMaxSupply && valFound.MaxSupply != msg.MaxSupply {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "cannot change maxsupply")
	}
	if !valFound.CanChangeMaxSupply && msg.CanChangeMaxSupply {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot revert change maxsupply flag")
	}
	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             valFound.Ticker,
		Precision:          valFound.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		Supply:             valFound.Supply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}
