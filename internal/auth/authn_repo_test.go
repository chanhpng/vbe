package auth_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chanhpng/vbe/internal/auth"
	"github.com/chanhpng/vbe/internal/repotesting"
	"github.com/chanhpng/vbe/internal/user"
	"github.com/chanhpng/vbe/repo"
)

func TestRepositoryAuthenticator(t *testing.T) {
	a := auth.AuthenticateRepositoryUsers()
	ctx, env := repotesting.NewEnvironment(t, repotesting.FormatNotImportant)

	require.NoError(t, repo.WriteSession(ctx, env.Repository, repo.WriteSessionOptions{},
		func(ctx context.Context, w repo.RepositoryWriter) error {
			for _, tc := range []struct {
				profile  *user.Profile
				password string
			}{
				{
					profile: &user.Profile{
						Username:            "user1@host1",
						PasswordHashVersion: user.ScryptHashVersion,
					},
					password: "password1",
				},
				{
					profile: &user.Profile{
						Username:            "user2@host2",
						PasswordHashVersion: user.ScryptHashVersion,
					},
					password: "password2",
				},
				{
					profile: &user.Profile{
						Username: "user3@host3",
					},
					password: "password3",
				},
				{
					profile: &user.Profile{
						Username:            "user4@host4",
						PasswordHashVersion: user.Pbkdf2HashVersion,
					},
					password: "password4",
				},
			} {
				tc.profile.SetPassword(tc.password)
				err := user.SetUserProfile(ctx, w, tc.profile)
				if err != nil {
					return err
				}
			}
			return nil
		}))

	verifyRepoAuthenticator(ctx, t, a, env.Repository, "user1@host1", "password1", true)
	verifyRepoAuthenticator(ctx, t, a, env.Repository, "user1@host1", "password2", false)
	verifyRepoAuthenticator(ctx, t, a, env.Repository, "user1@host1", "password11", false)
	verifyRepoAuthenticator(ctx, t, a, env.Repository, "user1@host1a", "password1", false)
	verifyRepoAuthenticator(ctx, t, a, env.Repository, "user1@host1a", "password1a", false)

	// Test for password with KeyDerivationSet
	verifyRepoAuthenticator(ctx, t, a, env.Repository, "user2@host2", "password2", true)

	// Test for User with neither key derivation or PasswordHashVersion set
	verifyRepoAuthenticator(ctx, t, a, env.Repository, "user3@host3", "password3", false)

	// Test for PBKDF2 key derivation
	verifyRepoAuthenticator(ctx, t, a, env.Repository, "user4@host4", "password4", true)
}

func verifyRepoAuthenticator(ctx context.Context, t *testing.T, a auth.Authenticator, r repo.Repository, username, password string, want bool) {
	t.Helper()

	if got := a.IsValid(ctx, r, username, password); got != want {
		t.Errorf("invalid authenticator result for %v/%v: %v, want %v", username, password, got, want)
	}
}
