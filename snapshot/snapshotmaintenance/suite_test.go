package snapshotmaintenance_test

import (
	"testing"

	"github.com/chanhpng/vbe/internal/testutil"
	"github.com/chanhpng/vbe/repo/format"
)

type formatSpecificTestSuite struct {
	formatVersion format.Version
}

func TestFormatV1(t *testing.T) {
	testutil.RunAllTestsWithParam(t, &formatSpecificTestSuite{format.FormatVersion1})
}

func TestFormatV2(t *testing.T) {
	testutil.RunAllTestsWithParam(t, &formatSpecificTestSuite{format.FormatVersion2})
}

func TestFormatV3(t *testing.T) {
	testutil.RunAllTestsWithParam(t, &formatSpecificTestSuite{format.FormatVersion3})
}
