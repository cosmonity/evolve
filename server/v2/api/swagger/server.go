package swagger

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"

	serverv2 "go.cosmonity.xyz/evolve/server/v2"

	"cosmossdk.io/core/server"
	"cosmossdk.io/core/transaction"
	"cosmossdk.io/log"
)

var (
	_ serverv2.ServerComponent[transaction.Tx] = (*Server[transaction.Tx])(nil)
	_ serverv2.HasConfig                       = (*Server[transaction.Tx])(nil)
)

const ServerName = "swagger"

// Server represents a Swagger UI server
type Server[T transaction.Tx] struct {
	logger     log.Logger
	config     *Config
	cfgOptions []CfgOption

	server *http.Server
}

// New creates a new Swagger UI server
func New[T transaction.Tx](
	logger log.Logger,
	swaggerUI fs.FS,
	config server.ConfigMap,
	cfgOptions ...CfgOption,
) (*Server[T], error) {
	s := &Server[T]{
		logger:     logger.With(log.ModuleKey, ServerName),
		cfgOptions: cfgOptions,
	}

	serverCfg := s.Config().(*Config)
	if len(config) > 0 {
		if err := serverv2.UnmarshalSubConfig(config, s.Name(), &serverCfg); err != nil {
			return s, fmt.Errorf("failed to unmarshal config: %w", err)
		}
	}
	s.config = serverCfg

	mux := http.NewServeMux()
	mux.Handle("/swagger/", &swaggerHandler{
		swaggerFS: swaggerUI,
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	s.server = &http.Server{
		Addr:    s.config.Address,
		Handler: mux,
	}

	return s, nil
}

// Name returns the server's name
func (s *Server[T]) Name() string {
	return ServerName
}

// Config returns the server configuration
func (s *Server[T]) Config() any {
	if s.config == nil || s.config.Address == "" {
		cfg := DefaultConfig()
		// overwrite the default config with the provided options
		for _, opt := range s.cfgOptions {
			opt(cfg)
		}

		return cfg
	}

	return s.config
}

// Start starts the server
func (s *Server[T]) Start(ctx context.Context) error {
	if !s.config.Enable {
		s.logger.Info(fmt.Sprintf("%s server is disabled via config", s.Name()))
		return nil
	}

	s.logger.Info("starting swagger server...", "address", s.config.Address)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start swagger server: %w", err)
	}

	return nil
}

// Stop stops the server
func (s *Server[T]) Stop(ctx context.Context) error {
	if !s.config.Enable {
		return nil
	}

	s.logger.Info("stopping swagger server...", "address", s.config.Address)
	return s.server.Shutdown(ctx)
}