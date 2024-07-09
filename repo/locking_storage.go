package repo

import (
	"github.com/chanhpng/vbe/internal/epoch"
	"github.com/chanhpng/vbe/repo/content"
	"github.com/chanhpng/vbe/repo/content/indexblob"
	"github.com/chanhpng/vbe/repo/format"
)

// GetLockingStoragePrefixes Return all prefixes that may be maintained by Object Locking.
func GetLockingStoragePrefixes() []string {
	var prefixes []string
	// collect prefixes that need to be locked on put
	for _, prefix := range content.PackBlobIDPrefixes {
		prefixes = append(prefixes, string(prefix))
	}

	prefixes = append(prefixes, indexblob.V0IndexBlobPrefix, epoch.EpochManagerIndexUberPrefix, format.KopiaRepositoryBlobID,
		format.KopiaBlobCfgBlobID)

	return prefixes
}
