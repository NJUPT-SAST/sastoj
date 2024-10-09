package service

import (
	"context"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/biz"
)

func (s *AdminService) CreateProblem(ctx context.Context, request *pb.CreateProblemRequest) (*pb.CreateProblemReply, error) {
	rv, err := s.pc.CreateProblem(ctx, &biz.Problem{
		TypeId:     request.TypeId,
		Title:      request.Title,
		Content:    request.Content,
		Point:      request.Point,
		ContestId:  request.ContestId,
		Index:      request.Index,
		Visibility: request.Visibility,
		OwnerId:    request.OwnerId,
		Config:     request.Config,
		Metadata:   request.Metadata,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateProblemReply{
		Id: *rv,
	}, nil
}

func (s *AdminService) UpdateProblem(ctx context.Context, request *pb.UpdateProblemRequest) (*pb.UpdateProblemReply, error) {
	rv, err := s.pc.UpdateProblem(ctx, &biz.Problem{
		TypeId:     request.TypeId,
		Title:      request.Title,
		Content:    request.Content,
		Point:      request.Point,
		ContestId:  request.ContestId,
		Index:      request.Index,
		Visibility: request.Visibility,
		OwnerId:    request.OwnerId,
		Config:     request.Config,
		Metadata:   request.Metadata,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProblemReply{
		Success: rv,
	}, nil
}

func (s *AdminService) DeleteProblem(ctx context.Context, request *pb.DeleteProblemRequest) (*pb.DeleteProblemReply, error) {
	rv, err := s.pc.DeleteProblem(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProblemReply{
		Success: rv,
	}, nil
}

func (s *AdminService) GetProblem(ctx context.Context, request *pb.GetProblemRequest) (*pb.GetProblemReply, error) {
	rv, err := s.pc.FindProblem(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProblemReply{
		Id:          rv.Id,
		TypeId:      rv.TypeId,
		Title:       rv.Title,
		Content:     rv.Content,
		Point:       rv.Point,
		ContestId:   rv.ContestId,
		CaseVersion: rv.CaseVersion,
		Index:       rv.Index,
		Config:      rv.Config,
		OwnerId:     rv.OwnerId,
		Visibility:  rv.Visibility,
		Metadata:    rv.Metadata,
	}, nil
}

func (s *AdminService) ListProblem(ctx context.Context, request *pb.ListProblemRequest) (*pb.ListProblemReply, error) {
	rv, err := s.pc.ListProblem(ctx, request.Currency, request.Size)
	if err != nil {
		return nil, err
	}
	var list []*pb.ListProblemReply_Problem
	for _, each := range rv {
		r := pb.ListProblemReply_Problem{
			Id:          each.Id,
			TypeId:      each.TypeId,
			Title:       each.Title,
			Content:     each.Content,
			Point:       each.Point,
			ContestId:   each.ContestId,
			CaseVersion: each.CaseVersion,
			Index:       each.Index,
			Config:      each.Config,
			OwnerId:     each.OwnerId,
			Visibility:  each.Visibility,
			Metadata:    each.Metadata,
		}
		list = append(list, &r)
	}
	return &pb.ListProblemReply{
		Problems: list,
	}, nil
}
