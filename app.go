package go

import (
	"encoding/json"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice"
)

const appName = "nameservice"

var(
	// default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.nscli")
	
	// DefaultNodeHome sets the folder where the application data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.nsd")

	// NewBasicManager is in charge of setting up basic module elements
	ModuleBasic = module.NewBasicManager()

	// account permissions
	maccPerms = map[string][]string{}	
)

type nameServiceApp struct {
	*bam.BaseApp
}

func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {
		// First define the top level codec that will be shared by the different modules. Note: Codec will be explained later
		cdc: = MakeCodec()
		
		// BaseApp handles interactions with Tendermint through the ABCI protocol
		bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))
		
		var app = &nameServiceApp{
			BaseApp: bApp,
			cdc: cdc,
		}

		return app
}