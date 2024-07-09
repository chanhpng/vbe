package blobtesting

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/chanhpng/vbackup/internal/clock"
	"github.com/chanhpng/vbackup/internal/parallelwork"
	"github.com/chanhpng/vbackup/repo/blob"
)

// MinCleanupAge is the default cleanup age.
const MinCleanupAge = time.Hour

// CleanupOldData removes blobs older than provided time from storage using provided prefix.
func CleanupOldData(ctx context.Context, tb testing.TB, st blob.Storage, cleanupAge time.Duration) {
	tb.Helper()

	pq := parallelwork.NewQueue()

	now := clock.Now()

	_ = st.ListBlobs(ctx, "", func(it blob.Metadata) error {
		age := now.Sub(it.Timestamp)
		if age > cleanupAge {
			pq.EnqueueBack(ctx, func() error {
				tb.Logf("deleting %v", it.BlobID)

				return st.DeleteBlob(ctx, it.BlobID)
			})
		}
		return nil
	})

	require.NoError(tb, pq.Process(ctx, 16))
}
