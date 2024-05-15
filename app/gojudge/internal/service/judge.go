package service

import (
	"context"
	"sastoj/app/gojudge/internal/biz"

	pb "sastoj/api/sastoj/gojudge/judger/service/v1"
)

type JudgeService struct {
	pb.UnimplementedJudgeServer
	mu *biz.ManageUseCase
}

func NewJudgeService(useCase *biz.ManageUseCase) *JudgeService {
	return &JudgeService{
		mu: useCase,
	}
}

func (s *JudgeService) CreateOne(ctx context.Context, req *pb.CreateOneRequest) (*pb.CreateOneReply, error) {
	return s.mu.CreateOne(ctx, req)
}
func (s *JudgeService) FetchStatus(ctx context.Context, req *pb.FetchStatusRequest) (*pb.FetchStatusReply, error) {
	return s.mu.FetchStatus(ctx, req)
}
func (s *JudgeService) ModifyConfig(ctx context.Context, req *pb.ModifyConfigRequest) (*pb.ModifyConfigReply, error) {
	return s.mu.ModifyConfig(ctx, req)
}
func (s *JudgeService) DeleteOne(ctx context.Context, req *pb.DeleteOneRequest) (*pb.DeleteOneReply, error) {
	return s.mu.DeleteOne(ctx, req)
}
func (s *JudgeService) ListAll(ctx context.Context, req *pb.ListAllRequest) (*pb.ListAllReply, error) {
	return s.mu.ListAll(ctx, req)
}
