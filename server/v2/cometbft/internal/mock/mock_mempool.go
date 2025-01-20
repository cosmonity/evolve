package mock

import (
	"context"

	"go.cosmonity.xyz/evolve/server/v2/cometbft/mempool"

	"cosmossdk.io/core/transaction"
)

var _ mempool.Mempool[transaction.Tx] = (*MockMempool[transaction.Tx])(nil)

// MockMempool implements Mempool
// Used for testing instead of NoOpMempool
type MockMempool[T transaction.Tx] struct{}

func (MockMempool[T]) Insert(context.Context, T) error                 { return nil }
func (MockMempool[T]) Select(context.Context, []T) mempool.Iterator[T] { return nil }
func (MockMempool[T]) SelectBy(context.Context, []T, func(T) bool)     {}
func (MockMempool[T]) CountTx() int                                    { return 0 }
func (MockMempool[T]) Remove(T) error                                  { return nil }
