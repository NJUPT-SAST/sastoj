package limiter

import (
	"context"
	"errors"
	"fmt"
	"sastoj/pkg/middleware/auth"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

var blackList = map[int64]time.Time{}

func ApiLimiterMiddleware(apis map[string]time.Duration) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			trans, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("no Method Found")
			}
			operation := trans.Operation()
			t, exists := apis[operation]
			if !exists {
				return handler(ctx, req)
			}
			userID := ctx.Value("userInfo").(*auth.Claims).UserId
			prevTime, exists := blackList[userID]
			log.Debug(fmt.Sprintf("userID: %d, prevTime: %v, exists: %v", userID, prevTime, exists))
			if exists && time.Now().Before(prevTime) {
				log.Warn(fmt.Sprintf("user %d who tried to %s, was blocked by Api Limiter", userID, operation))
				return nil, errors.New("rate Limit Exceeded")
			}
			blackList[userID] = time.Now().Add(t)
			log.Debug(fmt.Sprintf("user %d who tried to %s, was allowed by Api Limiter", userID, operation))
			return handler(ctx, req)
		}
	}
}
