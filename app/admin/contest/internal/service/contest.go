package service

import (
	"context"
	"sastoj/app/admin/contest/internal/biz"

	pb "sastoj/api/sastoj/admin/contest"
)

type ContestService struct {
	pb.UnimplementedContestServer
	cs *biz.ContestUsecase
}

func NewContestService(cs *biz.ContestUsecase) *ContestService {
	return &ContestService{cs: cs}
}

func (s *ContestService) CreateContest(ctx context.Context, req *pb.CreateContestRequest) (*pb.CreateContestReply, error) {
	return &pb.CreateContestReply{}, nil
}
func (s *ContestService) UpdateContest(ctx context.Context, req *pb.UpdateContestRequest) (*pb.UpdateContestReply, error) {
	return &pb.UpdateContestReply{}, nil
}
func (s *ContestService) DeleteContest(ctx context.Context, req *pb.DeleteContestRequest) (*pb.DeleteContestReply, error) {
	return &pb.DeleteContestReply{}, nil
}
func (s *ContestService) GetContest(ctx context.Context, req *pb.GetContestRequest) (*pb.GetContestReply, error) {
	return &pb.GetContestReply{}, nil
}
func (s *ContestService) ListContest(ctx context.Context, req *pb.ListContestRequest) (*pb.ListContestReply, error) {
	return &pb.ListContestReply{}, nil
}
