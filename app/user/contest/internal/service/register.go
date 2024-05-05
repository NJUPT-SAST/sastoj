package service

import (
	"context"
	"github.com/google/uuid"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/contest/internal/biz"
)

func (s *UserContestService) RegisterGateway(ctx context.Context, req *pb.RegisterGatewayRequest) (*pb.RegisterGatewayReply, error) {
	ret := uuid.New().String()
	err := s.registerUc.RegisterGateway(ctx, &biz.Register{
		UUID:     ret,
		Endpoint: req.Endpoint,
		Secret:   req.Secret,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterGatewayReply{
		Token: ret,
	}, nil
}
