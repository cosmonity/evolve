package serverv2_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	serverv2 "go.cosmonity.xyz/evolve/server/v2"
	"go.cosmonity.xyz/evolve/server/v2/api/grpc"
	"go.cosmonity.xyz/evolve/server/v2/store"
	"go.cosmonity.xyz/evolve/store/v2/root"
)

func TestReadConfig(t *testing.T) {
	currentDir, err := os.Getwd()
	require.NoError(t, err)
	configPath := filepath.Join(currentDir, "testdata")

	v, err := serverv2.ReadConfig(configPath)
	require.NoError(t, err)

	require.Equal(t, v.GetString(grpc.FlagAddress), grpc.DefaultConfig().Address)
	require.Equal(t, v.GetString(store.FlagAppDBBackend), root.DefaultConfig().AppDBBackend)
}

func TestUnmarshalSubConfig(t *testing.T) {
	currentDir, err := os.Getwd()
	require.NoError(t, err)
	configPath := filepath.Join(currentDir, "testdata")

	v, err := serverv2.ReadConfig(configPath)
	require.NoError(t, err)
	cfg := v.AllSettings()

	grpcConfig := grpc.DefaultConfig()
	err = serverv2.UnmarshalSubConfig(cfg, "grpc", &grpcConfig)
	require.NoError(t, err)

	require.True(t, grpc.DefaultConfig().Enable)
	require.False(t, grpcConfig.Enable)

	storeConfig := root.Config{}
	err = serverv2.UnmarshalSubConfig(cfg, "store", &storeConfig)
	require.NoError(t, err)
	require.Equal(t, *root.DefaultConfig(), storeConfig)
}
