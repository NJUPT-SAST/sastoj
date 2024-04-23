package service

import (
	"context"

	pb "sastoj/api/sastoj/user/gateway/service/v1"
)

func (s *GatewayService) Submit(ctx context.Context, req *pb.SubmitRequest) (*pb.SubmitReply, error) {
	return &pb.SubmitReply{}, nil
}
func (s *GatewayService) SelfTest(ctx context.Context, req *pb.SelfTestRequest) (*pb.SelfTestReply, error) {
	return &pb.SelfTestReply{}, nil
}
func (s *GatewayService) GetSubmissions(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	return &pb.GetSubmissionReply{}, nil
}
func (s *GatewayService) ListRanking(ctx context.Context, req *pb.ListRankingRequest) (*pb.ListRankingReply, error) {
	return &pb.ListRankingReply{}, nil
}
