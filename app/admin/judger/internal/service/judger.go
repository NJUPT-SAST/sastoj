package service

import (
	"context"
	"sastoj/app/admin/judger/internal/biz"

	pb "sastoj/api/sastoj/admin/judger/service/v1"
)

type JudgerService struct {
	pb.UnimplementedJudgerServer
	ju *biz.JudgerUsecase
}

func NewJudgerService(judger *biz.JudgerUsecase) *JudgerService {
	return &JudgerService{
		ju: judger,
	}
}

func (s *JudgerService) UpdateJudger(ctx context.Context, req *pb.UpdateJudgerRequest) (*pb.UpdateJudgerReply, error) {
	err := s.ju.UpdateJudger(ctx, req.ProblemId, req.JudgerGroupIds)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateJudgerReply{
		Success: true,
	}, nil
}
func (s *JudgerService) GetJudger(ctx context.Context, req *pb.GetJudgerRequest) (*pb.GetJudgerReply, error) {
	res, err := s.ju.GetJudger(ctx, req.ProblemId)
	if err != nil {
		return nil, err
	}
	var list []*pb.GetJudgerReply_Group
	for _, v := range res.Groups {
		list = append(list, &pb.GetJudgerReply_Group{
			Id:   v.ID,
			Name: v.GroupName,
		})
	}
	return &pb.GetJudgerReply{
		Groups: list,
	}, nil
}
