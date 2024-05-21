package service

import (
	"context"
	pb "sastoj/api/sastoj/gojudge/judger/service/v1"
	"sastoj/app/gojudge/internal/biz"
)

type JudgeService struct {
	pb.UnimplementedJudgeServer
	judge *biz.Judge
}

func NewJudgeService(judge *biz.Judge) *JudgeService {
	return &JudgeService{judge: judge}
}

func (s *JudgeService) ListAllStatus(ctx context.Context, req *pb.ListAllStatusRequest) (*pb.ListAllStatusReply, error) {
	var status []*pb.Status
	for _, judge := range s.judge.Gojudge {
		status = append(status, &pb.Status{
			Type:     pb.JudgeType_gojudge,
			Endpoint: judge.Endpoint,
		})
	}
	return &pb.ListAllStatusReply{
		Status: status,
	}, nil

}
