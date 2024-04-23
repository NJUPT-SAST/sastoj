package util

import (
	"context"
	"crypto/tls"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/gateway/internal/conf"
)

func Sync(ctx context.Context, conf *conf.Client) {
	tlsConfig := tls.Config{
		ServerName: conf.Grpc.Addr,
	}
	conn, err := grpc.Dial(
		ctx,
		grpc.WithEndpoint(conf.Grpc.Endpoint),
		grpc.WithTLSConfig(&tlsConfig),
	)
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewContestServiceClient(conn)
	// TODO: get today contests
	_, err = client.ListContest(ctx, &pb.ListContestRequest{})
	if err != nil {
		return
	}

}
