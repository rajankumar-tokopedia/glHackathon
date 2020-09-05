package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/rajankumar549/glHackathon/src/apperror"
	"github.com/rajankumar549/glHackathon/src/constants"
	"github.com/rajankumar549/glHackathon/src/interfaces/server"
	"github.com/rajankumar549/glHackathon/src/usecases/auth"
)

func AuthRequest(ctx context.Context, next []server.HttpMiddleware, r *http.Request, p server.HttpParams) (interface{}, error) {
	if strings.Index(r.URL.Path, "auth") == -1 {
		return next[0](ctx, next[1:], r, p)
	}

	token := r.Header.Get(constants.AUTH_TOKEN)
	isAuthenticated, userInfo := auth.IsAuthenticated(ctx, token)
	if !isAuthenticated {
		return nil, apperror.ForbiddenError("UN_AUTH", "Unauthorized request")
	}

	ctx = context.WithValue(ctx, constants.CTX_GROUP_ID, userInfo.GroupId)
	ctx = context.WithValue(ctx, constants.CTX_USER_ID, userInfo.UserId)
	ctx = context.WithValue(ctx, constants.CTX_USER_INFO, userInfo)
	return next[0](ctx, next[1:], r, p)
}
