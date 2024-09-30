package util

import (
	"context"
	"sastoj/pkg/middleware/auth"
)

func GetUserInfoFromCtx(ctx context.Context) *auth.Claims {
	return ctx.Value("userInfo").(*auth.Claims)
}
