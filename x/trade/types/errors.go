package types

// DONTCOVER

import (
	sdkErrors "cosmossdk.io/errors"
)

// x/trade module sentinel errors
var (
	ErrTradeCreatedSuccessfully   = sdkErrors.Register(ModuleName, 1000, "trade created successfully")
	ErrTradeProcessedSuccessfully = sdkErrors.Register(ModuleName, 1001, "trade processed successfully")
	ErrSample                     = sdkErrors.Register(ModuleName, 1100, "sample error ")
	ErrInvalidTradeType           = sdkErrors.Register(ModuleName, 1101, "invalid trade type ")
	ErrInvalidTradeData           = sdkErrors.Register(ModuleName, 1102, "invalid trade data ")
	ErrInvalidTradePrice          = sdkErrors.Register(ModuleName, 1103, "invalid trade price ")
	ErrInvalidTradeQuantity       = sdkErrors.Register(ModuleName, 1104, "invalid trade quantity, max quantity is 9223372036854775807")
	ErrBurnCoins                  = sdkErrors.Register(ModuleName, 1105, "failed trade burn coins ")
	ErrMintCoins                  = sdkErrors.Register(ModuleName, 1106, "failed trade mint coins ")
	ErrInvalidReceiverAddress     = sdkErrors.Register(ModuleName, 1107, "invalid receiver address ")
	ErrModuleToAccountSendCoins   = sdkErrors.Register(ModuleName, 1108, "failed trade module to account send coin error ")
	ErrAccountToModuleSendCoins   = sdkErrors.Register(ModuleName, 1109, "failed trade account to module send coin error ")
	ErrInvalidCreator             = sdkErrors.Register(ModuleName, 1110, "invalid creator address ")
	ErrInvalidChecker             = sdkErrors.Register(ModuleName, 1111, "invalid checker address ")
	ErrInvalidProcessType         = sdkErrors.Register(ModuleName, 1112, "invalid process type ")
	ErrInvalidTradeIndex          = sdkErrors.Register(ModuleName, 1113, "invalid trade index ")
	ErrTradeStatusCompleted       = sdkErrors.Register(ModuleName, 1114, "trade is already completed ")
	ErrCheckerMustBeDifferent     = sdkErrors.Register(ModuleName, 1115, "checker must be different than maker ")
	ErrTradeStatusCanceled        = sdkErrors.Register(ModuleName, 1116, "trade is already canceled ")
	ErrTradeStatusRejected        = sdkErrors.Register(ModuleName, 1117, "trade is already rejected ")
	ErrInvalidMakerPermission     = sdkErrors.Register(ModuleName, 1118, "invalid maker permission ")
	ErrInvalidCheckerPermission   = sdkErrors.Register(ModuleName, 1119, "invalid checker permission ")
	ErrInvalidDateFormat          = sdkErrors.Register(ModuleName, 1120, "invalid date format ")
	ErrInvalidCreatorAddress      = sdkErrors.Register(ModuleName, 1121, "invalid creator address ")
	ErrInvalidCoinDenom           = sdkErrors.Register(ModuleName, 1122, "invalid coin denom ")
	ErrInvalidPath                = sdkErrors.Register(ModuleName, 1123, "invalid Path ")
	ErrInvalidStatus              = sdkErrors.Register(ModuleName, 1124, "invalid status ")
	ErrInvalidTradeDataObject     = sdkErrors.Register(ModuleName, 1125, "invalid trade data json")
	ErrInvalidTradeDataJSON       = sdkErrors.Register(ModuleName, 1126, "invalid trade data object")
	ErrAssetHolderID              = sdkErrors.Register(ModuleName, 1127, "invalid asset holder id")
	ErrAssetID                    = sdkErrors.Register(ModuleName, 1128, "invalid asset id")
	ErrTradeRequestID             = sdkErrors.Register(ModuleName, 1129, "invalid trade request id")
	ErrTradeValue                 = sdkErrors.Register(ModuleName, 1130, "invalid trade value")
	ErrTradeCurrency              = sdkErrors.Register(ModuleName, 1131, "invalid trade currency")
	ErrTradeExchange              = sdkErrors.Register(ModuleName, 1132, "invalid trade exchange")
	ErrTradeFundName              = sdkErrors.Register(ModuleName, 1133, "invalid trade fundName")
	ErrTradeIssuer                = sdkErrors.Register(ModuleName, 1134, "invalid trade issuer")
	ErrTradeNoShares              = sdkErrors.Register(ModuleName, 1135, "invalid trade number of shares")
	ErrTradeSegment               = sdkErrors.Register(ModuleName, 1136, "invalid trade segment")
	ErrTradeSharePrice            = sdkErrors.Register(ModuleName, 1137, "invalid trade share price")
	ErrTradeTicker                = sdkErrors.Register(ModuleName, 1138, "invalid trade ticker")
	ErrTradeFee                   = sdkErrors.Register(ModuleName, 1139, "invalid trade fee")
	ErrTradeNetPrice              = sdkErrors.Register(ModuleName, 1140, "invalid trade trade net price")
	ErrTradeNetValue              = sdkErrors.Register(ModuleName, 1141, "invalid trade net value")
)
