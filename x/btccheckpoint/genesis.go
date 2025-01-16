package btccheckpoint

import (
	"context"
	"log"

	"github.com/babylonlabs-io/babylon/x/btccheckpoint/keeper"
	"github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
)

// InitGenesis initializes the module's state using the provided genesis state.
// If an error occurs while setting parameters, it logs the error and panics.
func InitGenesis(ctx context.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set module parameters
	if err := k.SetParams(ctx, genState.Params); err != nil {
		// Log the error before panicking
		log.Printf("Error initializing Genesis: %v", err)
		panic(err)
	}
}

// ExportGenesis exports the current module state as GenesisState.
func ExportGenesis(ctx context.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	return genesis
}
