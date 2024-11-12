package service

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
)

func (s *AdminService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	rv, err := s.gc.CreateGroup(ctx, &biz.Group{Name: req.Name, ManageId: req.Manage, ContestsId: req.Contests, ProblemsId: req.Problems})
	if err != nil {
		return nil, err
	}
	return &pb.CreateGroupReply{
		Id: rv.Id,
	}, nil
}
func (s *AdminService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*pb.UpdateGroupReply, error) {
	rv, err := s.gc.UpdateGroup(ctx, &biz.Group{
		Id:         req.Id,
		Name:       req.Name,
		ManageId:   req.Manage,
		ContestsId: req.Contests,
		ProblemsId: req.Problems,
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
			return nil, pb.ErrorGroupNotFound("Group with specified ID not found")
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
			return nil, pb.ErrorGroupNotFound("Group with specified ID not found")
		}
		return nil, err
	}
	var manage []*pb.GetContestReply
	var contests []*pb.GetContestReply
	var problems []*pb.Problem
	for _, m := range rv.Manage {
		manage = append(manage, &pb.GetContestReply{
			Id:          m.Id,
			Title:       m.Title,
			Description: m.Description,
			Status:      m.Status,
			Type:        m.Type,
			StartTime:   timestamppb.New(m.StartTime),
			EndTime:     timestamppb.New(m.EndTime),
			Language:    m.Language,
			ExtraTime:   m.ExtraTime,
			CreateTime:  timestamppb.New(m.CreateTime),
		})
	}
	for _, c := range rv.Contests {
		contests = append(contests, &pb.GetContestReply{
			Id:          c.Id,
			Title:       c.Title,
			Description: c.Description,
			Status:      c.Status,
			Type:        c.Type,
			StartTime:   timestamppb.New(c.StartTime),
			EndTime:     timestamppb.New(c.EndTime),
			Language:    c.Language,
			ExtraTime:   c.ExtraTime,
			CreateTime:  timestamppb.New(c.CreateTime),
		})
	}
	for _, p := range rv.Problems {
		problems = append(problems, &pb.Problem{
			Id:          p.ID,
			TypeId:      p.TypeID,
			Title:       p.Title,
			Content:     p.Content,
			Point:       p.Point,
			ContestId:   p.ContestID,
			CaseVersion: p.CaseVersion,
			Index:       p.Index,
		})
	}
	return &pb.GetGroupReply{
		Id:       rv.Id,
		Name:     rv.Name,
		Manage:   manage,
		Contests: contests,
		Problems: problems,
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
