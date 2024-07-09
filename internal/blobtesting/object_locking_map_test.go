package blobtesting

import (
	"testing"
	"time"

	"github.com/chanhpng/vbe/internal/testlogging"
	"github.com/chanhpng/vbe/repo/blob"
)

func TestObjectLockingStorage(t *testing.T) {
	r := NewVersionedMapStorage(nil)
	if r == nil {
		t.Errorf("unexpected result: %v", r)
	}

	VerifyStorage(testlogging.Context(t), t, r, blob.PutOptions{
		RetentionMode:   blob.Governance,
		RetentionPeriod: 24 * time.Hour,
	})
}
