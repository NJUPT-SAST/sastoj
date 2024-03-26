package service

import (
	"context"

	pb "sastoj/api/sastoj/user/contest/service/v1"
)

type ContestServiceService struct {
	pb.UnimplementedContestServiceServer
}

func NewContestServiceService() *ContestServiceService {
	return &ContestServiceService{}
}

func (s *ContestServiceService) ListContest(ctx context.Context, req *pb.ListContestRequest) (*pb.ListContestReply, error) {
	return &pb.ListContestReply{}, nil
}
func (s *ContestServiceService) JoinContest(ctx context.Context, req *pb.JoinContestRequest) (*pb.JoinContestReply, error) {
	return &pb.JoinContestReply{}, nil
}
func (s *ContestServiceService) GetProblem(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemReply, error) {
	return &pb.GetProblemReply{}, nil
}
func (s *ContestServiceService) SubmitProblem(ctx context.Context, req *pb.SubmitProblemRequest) (*pb.SubmitProblemReply, error) {
	return &pb.SubmitProblemReply{}, nil
}
func (s *ContestServiceService) PretestProblem(ctx context.Context, req *pb.PretestProblemRequest) (*pb.PretestProblemReply, error) {
	return &pb.PretestProblemReply{}, nil
}
func (s *ContestServiceService) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	return &pb.GetSubmissionReply{}, nil
}
func (s *ContestServiceService) GetRanking(ctx context.Context, req *pb.GetRankingRequest) (*pb.GetRankingReply, error) {
	return &pb.GetRankingReply{}, nil
}
