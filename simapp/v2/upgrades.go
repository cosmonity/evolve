package simapp

import (
	"context"

	"go.cosmonity.xyz/evolve/runtime/v2"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	upgradetypes "cosmossdk.io/x/upgrade/types"
)

// UpgradeName defines the on-chain upgrade name for the sample SimApp upgrade
// from v0.52.x to v2
//
// NOTE: This upgrade defines a reference implementation of what an upgrade
// could look like when an application is migrating from Cosmos SDK version
// v0.52.x to v2.
const UpgradeName = "v052-to-v2"

func (app *SimApp[T]) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(
		UpgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, fromVM appmodule.VersionMap) (appmodule.VersionMap, error) {
			return app.ModuleManager().RunMigrations(ctx, fromVM)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{
				"stf",
			},
		}

		app.SetStoreLoader(runtime.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
