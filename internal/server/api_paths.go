package server

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/chanhpng/vbe/internal/ospath"
	"github.com/chanhpng/vbe/internal/serverapi"
	"github.com/chanhpng/vbe/snapshot"
)

func handlePathResolve(ctx context.Context, rc requestContext) (interface{}, *apiError) {
	var req serverapi.ResolvePathRequest

	if err := json.Unmarshal(rc.body, &req); err != nil {
		return nil, requestError(serverapi.ErrorMalformedRequest, "malformed request body")
	}

	return &serverapi.ResolvePathResponse{
		SourceInfo: snapshot.SourceInfo{
			Path:     filepath.Clean(ospath.ResolveUserFriendlyPath(req.Path, true)),
			Host:     rc.rep.ClientOptions().Hostname,
			UserName: rc.rep.ClientOptions().Username,
		},
	}, nil
}
