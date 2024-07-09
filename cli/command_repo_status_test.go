package cli_test

import (
	"testing"

	"github.com/chanhpng/vbe/cli"
	"github.com/chanhpng/vbe/internal/testutil"
	"github.com/chanhpng/vbe/tests/testenv"
)

func TestRepoStatusJSON(t *testing.T) {
	t.Parallel()

	e := testenv.NewCLITest(t, testenv.RepoFormatNotImportant, testenv.NewInProcRunner(t))

	var rs cli.RepositoryStatus

	e.RunAndExpectSuccess(t, "repo", "create", "filesystem", "--path", e.RepoDir)
	defer e.RunAndExpectSuccess(t, "repo", "disconnect")

	e.RunAndExpectSuccess(t, "repo", "status")
	testutil.MustParseJSONLines(t, e.RunAndExpectSuccess(t, "repo", "status", "--json"), &rs)
}
