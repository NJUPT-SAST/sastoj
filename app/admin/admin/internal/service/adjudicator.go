package service

import (
	"context"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
)

func (s *AdminService) UpdateAdjudicator(ctx context.Context, req *pb.UpdateAdjudicatorRequest) (*pb.UpdateAdjudicatorReply, error) {
	err := s.jrc.UpdateAdjudicator(ctx, req.ProblemId, req.AdjudicatorGroupIds)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAdjudicatorReply{
		Success: true,
	}, nil
}
func (s *AdminService) GetAdjudicator(ctx context.Context, req *pb.GetAdjudicatorRequest) (*pb.GetAdjudicatorReply, error) {
	res, err := s.jrc.GetAdjudicator(ctx, req.ProblemId)
	if err != nil {
		return nil, err
	}
	var list []*pb.GetAdjudicatorReply_Group
	for _, v := range res.Groups {
		list = append(list, &pb.GetAdjudicatorReply_Group{
			Id:   v.ID,
			Name: v.GroupName,
		})
	}
	return &pb.GetAdjudicatorReply{
		Groups: list,
	}, nil
}
