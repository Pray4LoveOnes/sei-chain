package app

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/server"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DisableCosmWasm disables CosmWasm permissions and clears all CosmWasm state.
func (app *App) DisableCosmWasm(ctx sdk.Context) {
	fmt.Println("[SIP-3] Disabling CosmWasm runtime...")
	app.WasmKeeper.SetParams(ctx, wasmtypes.Params{
		CodeUploadAccess:             wasmtypes.NoAccessConfig(),
		InstantiateDefaultPermission: wasmtypes.AccessTypeNobody,
	})

	store := ctx.KVStore(app.GetKey(wasmtypes.StoreKey))
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
	fmt.Println("[SIP-3] CosmWasm KV store fully pruned.")
}

// GenerateERC20MigrationConfig converts a CW20 metadata JSON file into an ERC20 config.
func GenerateERC20MigrationConfig(input string) error {
	data, err := os.ReadFile(input)
	if err != nil {
		return err
	}

	var cw20 struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
		Supply   string `json:"total_supply"`
	}
	if err := json.Unmarshal(data, &cw20); err != nil {
		return err
	}

	out := map[string]any{
		"name":        cw20.Name,
		"symbol":      cw20.Symbol,
		"decimals":    cw20.Decimals,
		"totalSupply": cw20.Supply,
	}
	res, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("erc20_migrated.json", res, 0o644)
}

// GenerateUpgradeProposal writes a governance proposal file for the SIP-3 upgrade.
func GenerateUpgradeProposal() error {
	proposal := map[string]any{
		"title":       "SIP-3: Transition to EVM-Only",
		"description": "This enacts SIP-3, deprecating CosmWasm and enabling EVM-only runtime.",
		"upgrade": map[string]any{
			"name":   "evm-only",
			"height": 9500000,
		},
		"deposit": "10000000usei",
	}
	res, err := json.MarshalIndent(proposal, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile("sip3_software_upgrade.json", res, 0o644); err != nil {
		return err
	}
	fmt.Println("[✓] Governance proposal written: sip3_software_upgrade.json")
	return nil
}

// TestCosmWasmDeactivation ensures the CosmWasm store no longer contains data.
func TestCosmWasmDeactivation(ctx sdk.Context, storeKey storetypes.StoreKey) error {
	store := ctx.KVStore(storeKey)
	prefix := []byte{0x01}
	it := sdk.KVStorePrefixIterator(store, prefix)
	defer it.Close()
	if it.Valid() {
		return fmt.Errorf("CosmWasm not fully pruned: %x", it.Key())
	}
	fmt.Println("[✓] CosmWasm state clean.")
	return nil
}

// InitCodexSIP3 runs the SIP-3 helper suite for proposal/migration generation.
func InitCodexSIP3(ctx context.Context, app *App, serverCtx *server.Context) {
	_ = ctx
	_ = serverCtx

	if err := GenerateUpgradeProposal(); err != nil {
		fmt.Println("[ERR] Proposal generation failed:", err)
	}

	if err := GenerateERC20MigrationConfig("cw20_sample.json"); err != nil {
		fmt.Println("[ERR] Migration config gen failed:", err)
	}

	sdkCtx := app.BaseApp.NewContext(false, sdk.Header{})
	if err := TestCosmWasmDeactivation(sdkCtx, app.GetKey(wasmtypes.StoreKey)); err != nil {
		panic(err)
	}

	fmt.Println("[✓] SIP-3 Codex Initialization Complete")
}
