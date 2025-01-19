package mock

import "go.cosmonity.xyz/evolve/store/v2"

// StateCommitter is a mock of store.Committer
type StateCommitter interface {
	store.Committer
	store.Pruner
	store.PausablePruner
	store.UpgradeableStore
	store.VersionedReader
	store.UpgradableDatabase
}
