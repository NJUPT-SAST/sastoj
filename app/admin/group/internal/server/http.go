package server

import (
	v1 "sastoj/api/sastoj/admin/group/service/v1"
	"sastoj/app/admin/group/internal/conf"
	"sastoj/app/admin/group/internal/service"
	"sastoj/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, group *service.GroupService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			validate.Validator(),
			recovery.Recovery(),
			auth.Auth(c.JwtSecret, auth.AdminGroup, nil),
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
	// 错误解码器
	opts = append(opts, http.ErrorEncoder(ErrorEncoder))
	// 返回参数解码器
	opts = append(opts, http.ResponseEncoder(ResponseEncoder))
	srv := http.NewServer(opts...)
	v1.RegisterGroupHTTPServer(srv, group)
	return srv
}
