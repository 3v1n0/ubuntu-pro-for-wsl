// Package proservices is in charge of managing the GRPC services and all business-logic side.
package proservices

import (
	"context"
	"os"
	"path/filepath"

	agent_api "github.com/canonical/ubuntu-pro-for-windows/agentapi"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/consts"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/distroDB"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/grpc/interceptorschain"
	log "github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/grpc/logstreamer"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/proservices/ui"
	"github.com/canonical/ubuntu-pro-for-windows/windows-agent/internal/proservices/wslinstance"
	"google.golang.org/grpc"
)

// Manager is the orchestrator of GRPC API services and business logic.
type Manager struct {
	uiService          ui.Service
	wslInstanceService wslinstance.Service
}

// options are the configurable functional options for the daemon.
type options struct {
	cacheDir string
}

// Option is the function signature we are passing to tweak the daemon creation.
type Option func(*options) error

// WithCacheDir overrides the cache directory used in the daemon.
func WithCacheDir(cachedir string) func(o *options) error {
	return func(o *options) error {
		if cachedir != "" {
			o.cacheDir = cachedir
		}
		return nil
	}
}

// New returns a new GRPC services manager.
// It instantiates both ui and wsl instance services.
func New(ctx context.Context, args ...Option) (s Manager, err error) {
	log.Debug(ctx, "Building new GRPC services manager")

	// Set default options.
	defaultUserCacheDir, err := os.UserCacheDir()
	if err != nil {
		return s, err
	}
	opts := options{
		cacheDir: filepath.Join(defaultUserCacheDir, consts.CacheBaseDirectory),
	}

	// Apply given options.
	for _, f := range args {
		if err := f(&opts); err != nil {
			return s, err
		}
	}

	log.Debugf(ctx, "Manager service cache directory: %s", opts.cacheDir)

	if err := os.MkdirAll(opts.cacheDir, 0750); err != nil {
		return s, err
	}

	db, err := distroDB.New(opts.cacheDir)
	if err != nil {
		return s, err
	}

	uiService, err := ui.New(ctx)
	if err != nil {
		return s, err
	}
	wslInstanceService, err := wslinstance.New(ctx, db)
	if err != nil {
		return s, err
	}
	return Manager{
		uiService:          uiService,
		wslInstanceService: wslInstanceService,
	}, nil
}

// RegisterGRPCServices returns a new grpc Server with the 2 api services attached to it.
// It also gets the correct middlewares hooked in.
func (m Manager) RegisterGRPCServices(ctx context.Context) *grpc.Server {
	log.Debug(ctx, "Registering GRPC services")

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(
		interceptorschain.StreamServer(
		/*log.StreamServerInterceptor(logrus.StandardLogger()),
		logconnections.StreamServerInterceptor(),*/
		)))
	agent_api.RegisterUIServer(grpcServer, m.uiService)
	agent_api.RegisterWSLInstanceServer(grpcServer, &m.wslInstanceService)

	return grpcServer
}
