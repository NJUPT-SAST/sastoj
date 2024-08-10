package server

import (
	"context"
	"fmt"
	"net"
	v1 "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/gateway/internal/conf"
	"sastoj/app/user/gateway/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, gateway *service.GatewayService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			AntiCheatMiddleware(),
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
	v1.RegisterContestHTTPServer(srv, gateway)
	return srv
}

func AntiCheatMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if header, ok := tr.(*http.Transport); ok {
					//var ip string
					//for _, ip = range strings.Split(header.Get("X-Forwarded-For"), ",") {
					//	if ip != "" {
					//		break
					//	}
					//	fmt.Printf("x-forwarded-for: %s\n", ip)
					//}
					//ip = strings.TrimSpace(header.Get("X-Real-Ip"))
					//fmt.Printf("x-real-ip: %s\n", ip)
					ip, _, _ := net.SplitHostPort(header.Request().RemoteAddr)
					fmt.Printf("remote-addr: %s\n", ip)
					ctx = context.WithValue(ctx, "ip", ip)
					defer func() {
						// Do something on exiting
					}()
				}
			}
			return handler(ctx, req)
		}
	}
}
