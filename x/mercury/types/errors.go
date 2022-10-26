package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mercury module sentinel errors
var (
	ErrProviderBadSigner                     = sdkerrors.Register(ModuleName, 2, "unauthorized: bad provider pubkey and signer")
	ErrProviderAlreadyExists                 = sdkerrors.Register(ModuleName, 3, "provider already exists")
	ErrInsufficientFunds                     = sdkerrors.Register(ModuleName, 4, "insufficient funds")
	ErrInvalidBond                           = sdkerrors.Register(ModuleName, 5, "invalid bond")
	ErrInvalidModProviderMetdataURI          = sdkerrors.Register(ModuleName, 6, "invalid mod provider metadata uri")
	ErrInvalidModProviderMaxContractDuration = sdkerrors.Register(ModuleName, 7, "invalid mod provider max contract duration")
	ErrInvalidModProviderMinContractDuration = sdkerrors.Register(ModuleName, 8, "invalid mod provider min contract duration")
	ErrInvalidModProviderStatus              = sdkerrors.Register(ModuleName, 9, "invalid mod provider bad provider status")
	ErrInvalidModProviderNoBond              = sdkerrors.Register(ModuleName, 10, "no bond")
	ErrDisabledHandler                       = sdkerrors.Register(ModuleName, 11, "disabled handler")
	ErrInvalidChain                          = sdkerrors.Register(ModuleName, 12, "invalid chain")
	ErrOpenContractDuration                  = sdkerrors.Register(ModuleName, 13, "invalid contract duration")
	ErrOpenContractMismatchRate              = sdkerrors.Register(ModuleName, 14, "mismatch contract rate")
	ErrOpenContractAlreadyOpen               = sdkerrors.Register(ModuleName, 15, "contract is already open")
	ErrOpenContractRate                      = sdkerrors.Register(ModuleName, 16, "invalid contract rate")
	ErrInvalidContractType                   = sdkerrors.Register(ModuleName, 17, "invalid contract type")
)
