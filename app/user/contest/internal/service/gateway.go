package service

import (
	"context"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/contest/internal/biz"

	"github.com/google/uuid"
)

func (s *ContestService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	ret := uuid.New().String()
	err := s.registerUc.RegisterGateway(ctx, &biz.Register{
		UUID:     ret,
		Endpoint: req.Endpoint,
		Secret:   req.Secret,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		Token: ret,
	}, nil
}
