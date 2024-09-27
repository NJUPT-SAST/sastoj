package server

import (
	"github.com/go-kratos/kratos/v2/middleware/validate"
	v1 "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/conf"
	"sastoj/app/admin/admin/internal/service"
	"sastoj/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, admin *service.AdminService, logger log.Logger) *http.Server {
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
	srv := http.NewServer(opts...)
	customRoute := srv.Route("/")
	customRoute.POST("/case/upload", func(ctx http.Context) error {
		return admin.UploadCases(ctx)
	})
	v1.RegisterAdminHTTPServer(srv, admin)
	return srv
}
