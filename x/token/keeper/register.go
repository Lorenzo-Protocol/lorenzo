package keeper

import (
	"github.com/ethereum/go-ethereum/common"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

// RegisterCoin creates a token pair for existing coin and deploy its erc20 contract.
// NOTE: the mapping between sdk.Coin metadata and erc20 contract compromise the rules as follows:
func (k Keeper) RegisterCoin(ctx sdk.Context, coinMetadata banktypes.Metadata) (*types.TokenPair, error) {
	// check if coin denom is already registered
	if k.IsRegisteredByDenom(ctx, coinMetadata.Base) {
		return nil, errorsmod.Wrapf(
			types.ErrTokenPairAlreadyExists, "coin denomination already registered: %s", coinMetadata.Name,
		)
	}

	// check if the coin exists by ensuring the supply is set
	if !k.bankKeeper.HasSupply(ctx, coinMetadata.Base) {
		return nil, errorsmod.Wrapf(
			errortypes.ErrInvalidCoins, "base denomination '%s' cannot have a supply of 0", coinMetadata.Base,
		)
	}

	addr, err := k.DeployERC20Contract(ctx, coinMetadata)
	if err != nil {
		return nil, errorsmod.Wrap(
			err, "failed to create wrapped coin denom metadata for ERC20",
		)
	}

	pair := types.NewTokenPair(addr, coinMetadata.Base, types.OWNER_MODULE)
	id := pair.GetID()
	k.SetTokenPair(ctx, pair)
	k.SetTokenPairIdByDenom(ctx, pair.Denom, id)
	k.SetTokenPairIdByERC20(ctx, common.HexToAddress(pair.ContractAddress), id)

	return &pair, nil
}

func (k Keeper) RegisterERC20(ctx sdk.Context, contract common.Address) (*types.TokenPair, error) {
	if k.IsRegisteredByERC20(ctx, contract) {
		return nil, errorsmod.Wrapf(
			types.ErrTokenPairAlreadyExists, "token ERC20 contract already registered: %s", contract.String(),
		)
	}

	metadata, err := k.CreateCoinMetadata(ctx, contract)
	if err != nil {
		return nil, errorsmod.Wrap(
			err, "failed to create wrapped coin denom metadata for ERC20",
		)
	}

	pair := types.NewTokenPair(contract, metadata.Name, types.OWNER_EXTERNAL)
	id := pair.GetID()
	k.SetTokenPair(ctx, pair)
	k.SetTokenPairIdByDenom(ctx, pair.Denom, id)
	k.SetTokenPairIdByERC20(ctx, common.HexToAddress(pair.ContractAddress), id)

	return &pair, nil
}

func (k Keeper) CreateCoinMetadata(ctx sdk.Context, contract common.Address) (*banktypes.Metadata, error) {
	erc20Data, err := k.QueryERC20Contract(ctx, contract)
	if err != nil {
		return nil, err
	}

	// check if base denom already exists
	denom := erc20Data.BaseDenom()
	if _, found := k.bankKeeper.GetDenomMetaData(ctx, denom); found {
		return nil, errorsmod.Wrap(
			types.ErrInternalTokenPair, "denom metadata already registered",
		)
	}

	// check if denom registered
	if k.IsRegisteredByDenom(ctx, denom) {
		return nil, errorsmod.Wrapf(
			types.ErrInternalTokenPair, "coin denomination already registered: %s", erc20Data.Name,
		)
	}

	metadata := banktypes.Metadata{
		Description: erc20Data.Description(),
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    denom,
				Exponent: 0,
			},
		},
		Base:    denom,
		Display: denom,
		Name:    denom,
		Symbol:  erc20Data.Symbol,
	}

	// NOTE: if decimals > 0, add a new denom unit after sanitizing its format.
	// On this time, display need to be changed to wanted denom.
	if erc20Data.Decimals > 0 {
		nameSanitized := erc20Data.SanitizedName()
		metadata.Display = nameSanitized
		metadata.DenomUnits = append(
			metadata.DenomUnits,
			&banktypes.DenomUnit{
				Denom:    nameSanitized,
				Exponent: uint32(erc20Data.Decimals),
			},
		)
	}

	if err := metadata.Validate(); err != nil {
		return nil, errorsmod.Wrapf(
			err, "ERC20 token data is invalid for contract %s", contract.String(),
		)
	}

	k.bankKeeper.SetDenomMetaData(ctx, metadata)

	return &metadata, nil
}

// ToggleConversion toggles the conversion of a token pair.
func (k Keeper) ToggleConversion(ctx sdk.Context, token string) (types.TokenPair, error) {
	id := k.GetTokenPairId(ctx, token)
	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "token '%s' not registered", token,
		)
	}

	pair.Enabled = !pair.Enabled
	k.SetTokenPair(ctx, pair)

	return pair, nil
}
