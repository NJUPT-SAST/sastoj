package service

import (
	"context"

	pb "sastoj/api/sastoj/user/gateway/service/v1"
)

func (s *GatewayService) GetContests(ctx context.Context, req *pb.GetContestsRequest) (*pb.GetContestsReply, error) {
	return &pb.GetContestsReply{}, nil
}
func (s *GatewayService) JoinContest(ctx context.Context, req *pb.JoinContestRequest) (*pb.JoinContestReply, error) {
	return &pb.JoinContestReply{}, nil
}
