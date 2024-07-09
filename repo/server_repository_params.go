package repo

import (
	"github.com/chanhpng/vbe/internal/cache"
	"github.com/chanhpng/vbe/internal/metrics"
	"github.com/chanhpng/vbe/repo/format"
	"github.com/chanhpng/vbe/repo/hashing"
)

// immutableServerRepositoryParameters contains immutable parameters shared between HTTP and GRPC clients.
type immutableServerRepositoryParameters struct {
	h               hashing.HashFunc
	objectFormat    format.ObjectFormat
	cliOpts         ClientOptions
	metricsRegistry *metrics.Registry
	contentCache    *cache.PersistentCache
	beforeFlush     []RepositoryWriterCallback

	*refCountedCloser
}

// Metrics provides access to the metrics registry.
func (r *immutableServerRepositoryParameters) Metrics() *metrics.Registry {
	return r.metricsRegistry
}

func (r *immutableServerRepositoryParameters) ClientOptions() ClientOptions {
	return r.cliOpts
}
