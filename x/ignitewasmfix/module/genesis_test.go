package ignitewasmfix_test

import (
	"testing"

	keepertest "ignite-wasm-fix/testutil/keeper"
	"ignite-wasm-fix/testutil/nullify"
	ignitewasmfix "ignite-wasm-fix/x/ignitewasmfix/module"
	"ignite-wasm-fix/x/ignitewasmfix/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitewasmfixKeeper(t)
	ignitewasmfix.InitGenesis(ctx, k, genesisState)
	got := ignitewasmfix.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
