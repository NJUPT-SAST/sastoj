package service

import (
	"context"

	pb "sastoj/api/sastoj/user/gateway/service/v1"
)

func (s *GatewayService) GetProblems(ctx context.Context, req *pb.GetProblemsRequest) (*pb.GetProblemsReply, error) {
	return &pb.GetProblemsReply{}, nil
}
func (s *GatewayService) GetProblem(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemReply, error) {
	return &pb.GetProblemReply{}, nil
}
