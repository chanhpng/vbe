package snapshotfs_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chanhpng/vbe/fs"
	"github.com/chanhpng/vbe/fs/virtualfs"
	"github.com/chanhpng/vbe/internal/mockfs"
	"github.com/chanhpng/vbe/internal/testlogging"
	"github.com/chanhpng/vbe/snapshot"
	"github.com/chanhpng/vbe/snapshot/policy"
	"github.com/chanhpng/vbe/snapshot/snapshotfs"
)

type fakeProgress struct {
	t                   *testing.T
	expectedFiles       int32
	expectedDirectories int32
	expectedErrors      int32
}

func (p *fakeProgress) Processing(context.Context, string) {}

func (p *fakeProgress) Error(context.Context, string, error, bool) {}

// +checklocksignore.
func (p *fakeProgress) Stats(
	ctx context.Context,
	s *snapshot.Stats,
	includedFiles, excludedFiles snapshotfs.SampleBuckets,
	excludedDirs []string,
	final bool,
) {
	if !final {
		return
	}

	assert.Equal(p.t, p.expectedErrors, s.ErrorCount)
	assert.Equal(p.t, p.expectedFiles, s.TotalFileCount)
	assert.Equal(p.t, p.expectedDirectories, s.TotalDirectoryCount)
}

func TestEstimate_SkipsStreamingDirectory(t *testing.T) {
	f := mockfs.NewFile("f1", []byte{1, 2, 3}, 0o777)

	rootDir := virtualfs.NewStaticDirectory("root", []fs.Entry{
		virtualfs.NewStreamingDirectory(
			"a-dir",
			fs.StaticIterator([]fs.Entry{f}, nil),
		),
	})

	policyTree := policy.BuildTree(nil, policy.DefaultPolicy)
	p := &fakeProgress{
		t:                   t,
		expectedFiles:       0,
		expectedDirectories: 2,
		expectedErrors:      0,
	}

	err := snapshotfs.Estimate(testlogging.Context(t), rootDir, policyTree, p, 1)
	require.NoError(t, err)
}
