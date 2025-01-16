package btccheckpoint_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/babylonlabs-io/babylon/x/btccheckpoint"
	"github.com/stretchr/testify/require"

	simapp "github.com/babylonlabs-io/babylon/app"
	"github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
)

func TestExportGenesis(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false)

	// Set default parameters
	if err := app.BtcCheckpointKeeper.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	// Export genesis state and verify
	genesisState := btccheckpoint.ExportGenesis(ctx, app.BtcCheckpointKeeper)
	require.Equal(t, genesisState.Params, types.DefaultParams())
}

func TestInitGenesis(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false)

	// Create a valid GenesisState
	genesisState := types.GenesisState{
		Params: types.Params{
			BtcConfirmationDepth:          888,
			CheckpointFinalizationTimeout: 999,
			CheckpointTag:                 types.DefaultCheckpointTag,
		},
	}

	// Initialize genesis state
	btccheckpoint.InitGenesis(ctx, app.BtcCheckpointKeeper, genesisState)

	// Verify the parameters were set correctly
	require.Equal(t, app.BtcCheckpointKeeper.GetParams(ctx).BtcConfirmationDepth, uint32(888))
	require.Equal(t, app.BtcCheckpointKeeper.GetParams(ctx).CheckpointFinalizationTimeout, uint32(999))
}

func TestInitGenesis_WithError(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false)

	// Create an invalid GenesisState to trigger an error
	genesisState := types.GenesisState{
		Params: types.Params{}, // Missing required parameters
	}

	// Capture log output for verification
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Check that panic occurs
	defer func() {
		if r := recover(); r != nil {
			// Verify the log message contains the expected error
			require.Contains(t, buf.String(), "Error initializing Genesis")
		}
	}()

	// Call InitGenesis
	btccheckpoint.InitGenesis(ctx, app.BtcCheckpointKeeper, genesisState)
}
