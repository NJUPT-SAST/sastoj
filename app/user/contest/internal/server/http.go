package server

import (
	v1 "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/contest/internal/conf"
	"sastoj/app/user/contest/internal/service"
	"sastoj/pkg/middleware/auth"
	"sastoj/pkg/middleware/limiter"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, contest *service.ContestService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			auth.Auth(c.JwtSecret, auth.UserGroup, nil),
			limiter.ApiLimiterMiddleware(apiLimiter),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterContestHTTPServer(srv, contest)
	return srv
}
