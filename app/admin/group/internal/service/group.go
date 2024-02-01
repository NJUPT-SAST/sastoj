package service

import (
	"context"
	pb "sastoj/api/sastoj/admin/group/service/v1"
	"sastoj/app/admin/group/internal/biz"
)

type GroupService struct {
	pb.UnimplementedGroupServer
	gc *biz.GroupUsecase
}

func NewGroupService(group *biz.GroupUsecase) *GroupService {
	return &GroupService{
		gc: group,
	}
}

func (s *GroupService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	rv, err := s.gc.CreateGroup(ctx, &biz.Group{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.CreateGroupReply{
		Id: rv.Id,
	}, nil
}
func (s *GroupService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*pb.UpdateGroupReply, error) {
	return &pb.UpdateGroupReply{}, nil
}
func (s *GroupService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupReply, error) {
	return &pb.DeleteGroupReply{}, nil
}
func (s *GroupService) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupReply, error) {
	rv, err := s.gc.GetGroup(ctx, &biz.Group{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupReply{
		Id:   rv.Id,
		Name: rv.Name,
	}, nil
}
func (s *GroupService) ListGroup(ctx context.Context, req *pb.ListGroupRequest) (*pb.ListGroupReply, error) {
	return &pb.ListGroupReply{}, nil
}
