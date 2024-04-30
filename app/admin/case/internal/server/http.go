package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "sastoj/api/sastoj/admin/case/service/v1"
	"sastoj/app/admin/case/internal/conf"
	"sastoj/app/admin/case/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, problemCase *service.CaseService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
		return problemCase.UploadCases(ctx)
	})
	//customRoute.POST("/case/upload", handler.UploadCasesHandler)
	//customRoute.POST("/case/upload", func(ctx http.Context) error {
	//
	//	//ids, err := problemCase.UploadCases(ctx)
	//	//byteIds := make([]byte, len(ids))
	//	//for i, val := range ids {
	//	//	byteIds[i] = byte(val)
	//	//}
	//	//ctx.Response().Write(byteIds)
	//
	//	ctx.Response().Write([]byte("dwsaewfsdaf"))
	//	ctx.Result(200, *UploadCasesReply)
	//	ctx.Response().Header().Set("Content-Type", "application/json")
	//	reply := out.(*UploadCasesReply)
	//	return ctx.Result(200, reply)
	//	return nil
	//})
	v1.RegisterCaseServiceHTTPServer(srv, problemCase)
	return srv
}
