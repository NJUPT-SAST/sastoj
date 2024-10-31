package service

import (
	"context"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/biz"
)

func (s *AdminService) CreateProblem(ctx context.Context, request *pb.CreateProblemRequest) (*pb.CreateProblemReply, error) {
	rv, err := s.pc.CreateProblem(ctx, &biz.Problem{
		TypeID:     request.TypeId,
		Title:      request.Title,
		Content:    request.Content,
		Point:      request.Point,
		ContestID:  request.ContestId,
		Index:      request.Index,
		Visibility: request.Visibility,
		OwnerID:    request.OwnerId,
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
		TypeID:     request.TypeId,
		Title:      request.Title,
		Content:    request.Content,
		Point:      request.Point,
		ContestID:  request.ContestId,
		Index:      request.Index,
		Visibility: request.Visibility,
		OwnerID:    request.OwnerId,
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
		Id:          rv.ID,
		TypeId:      rv.TypeID,
		Title:       rv.Title,
		Content:     rv.Content,
		Point:       rv.Point,
		ContestId:   rv.ContestID,
		CaseVersion: rv.CaseVersion,
		Index:       rv.Index,
		Config:      rv.Config,
		OwnerId:     rv.OwnerID,
		Visibility:  rv.Visibility,
		Metadata:    rv.Metadata,
	}, nil
}

func (s *AdminService) ListProblem(ctx context.Context, request *pb.ListProblemRequest) (*pb.ListProblemReply, error) {
	rv, err := s.pc.ListProblem(ctx, request.Current, request.Size)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.ListProblemReply_Problem, 0)
	for _, each := range rv {
		r := pb.ListProblemReply_Problem{
			Id:          each.ID,
			TypeId:      each.TypeID,
			Title:       each.Title,
			Content:     each.Content,
			Point:       each.Point,
			ContestId:   each.ContestID,
			CaseVersion: each.CaseVersion,
			Index:       each.Index,
			Config:      each.Config,
			OwnerId:     each.OwnerID,
			Visibility:  each.Visibility,
			Metadata:    each.Metadata,
		}
		list = append(list, &r)
	}
	return &pb.ListProblemReply{
		Problems: list,
	}, nil
}

func (s *AdminService) GetProblemTypes(ctx context.Context, _ *pb.GetProblemTypesRequest) (*pb.GetProblemTypesReply, error) {
	rv, err := s.pc.GetProblemTypes(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.GetProblemTypesReply_ProblemType, 0)
	for _, each := range rv {
		r := pb.GetProblemTypesReply_ProblemType{
			Id:          each.ID,
			Name:        each.DisplayName,
			Slug:        each.SlugName,
			Description: each.Description,
			Judge:       each.Judge,
		}
		list = append(list, &r)
	}
	return &pb.GetProblemTypesReply{
		Types: list,
	}, nil
}
