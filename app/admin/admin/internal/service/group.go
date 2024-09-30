package service

import (
	"context"
	"errors"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
)

func (s *AdminService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	rv, err := s.gc.CreateGroup(ctx, &biz.Group{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.CreateGroupReply{
		Id: rv.Id,
	}, nil
}
func (s *AdminService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*pb.UpdateGroupReply, error) {
	rv, err := s.gc.UpdateGroup(ctx, &biz.Group{
		Id:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateGroupReply{
		Success: rv,
	}, nil
}
func (s *AdminService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupReply, error) {
	rv, err := s.gc.DeleteGroup(ctx, req.Id)
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			return nil, pb.ErrorGroupNotFound("Group with specified Id not found")
		}
		return nil, err
	}
	return &pb.DeleteGroupReply{
		Success: rv,
	}, nil

}
func (s *AdminService) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupReply, error) {
	rv, err := s.gc.GetGroup(ctx, &biz.Group{Id: req.Id})
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			return nil, pb.ErrorGroupNotFound("Group with specified Id not found")
		}
		return nil, err
	}
	return &pb.GetGroupReply{
		Id:   rv.Id,
		Name: rv.Name,
	}, nil
}
func (s *AdminService) ListGroup(ctx context.Context, req *pb.ListGroupRequest) (*pb.ListGroupReply, error) {
	rv, err := s.gc.ListGroup(ctx, req.Current, req.Size)
	if err != nil {
		return nil, err
	}
	var list []*pb.ListGroupReply_Group
	for _, v := range rv {
		list = append(list, &pb.ListGroupReply_Group{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	return &pb.ListGroupReply{
		Groups: list,
	}, nil

}
