package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v5"
	v1 "sastoj/api/sastoj/admin/contest/service/v1"
	"sastoj/app/admin/contest/internal/conf"
	"sastoj/app/admin/contest/internal/service"
	"sastoj/pkg/middleware/auth"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.ContestService, logger log.Logger) *http.Server {

	apiMap := map[string]string{
		"UpdateContest": "admin",
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			auth.Auth(func(token *jwt.Token) (interface{}, error) {
				return []byte("sastoj"), nil
			}, apiMap),
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
	v1.RegisterContestHTTPServer(srv, greeter)
	return srv
}
