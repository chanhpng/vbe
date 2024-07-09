package server

import (
	"context"

	"github.com/chanhpng/vbe/internal/serverapi"
	"github.com/chanhpng/vbe/repo"
)

func handleCurrentUser(ctx context.Context, _ requestContext) (interface{}, *apiError) {
	return serverapi.CurrentUserResponse{
		Username: repo.GetDefaultUserName(ctx),
		Hostname: repo.GetDefaultHostName(ctx),
	}, nil
}
