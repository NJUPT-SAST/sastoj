package service

import (
	"context"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
)

func (s *AdminService) UpdateJudger(ctx context.Context, req *pb.UpdateJudgerRequest) (*pb.UpdateJudgerReply, error) {
	err := s.jrc.UpdateJudger(ctx, req.ProblemId, req.JudgerGroupIds)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateJudgerReply{
		Success: true,
	}, nil
}
func (s *AdminService) GetJudger(ctx context.Context, req *pb.GetJudgerRequest) (*pb.GetJudgerReply, error) {
	res, err := s.jrc.GetJudger(ctx, req.ProblemId)
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
