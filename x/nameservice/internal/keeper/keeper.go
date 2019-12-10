package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice/types"
)


type Keeper struct {
	CoinKeeper bank.Keeper

	storeKey sdk.StoreKey		// Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}