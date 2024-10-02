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
		Score:      int16(request.Score),
		ContestID:  request.ContestId,
		Index:      int16(request.Index),
		Visibility: request.Visibility,
		OwnerID:    request.OwnerId,
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
		Score:      int16(request.Score),
		ContestID:  request.ContestId,
		Index:      int16(request.Index),
		Visibility: request.Visibility,
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
		Score:       int32(rv.Score),
		ContestId:   rv.ContestID,
		CaseVersion: int32(rv.CaseVersion),
		Index:       int32(rv.Index),
		OwnerId:     rv.OwnerID,
		Visibility:  rv.Visibility,
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
			Id:          each.ID,
			TypeId:      each.TypeID,
			Title:       each.Title,
			Content:     each.Content,
			Score:       int32(each.Score),
			ContestId:   each.ContestID,
			CaseVersion: int32(each.CaseVersion),
			Index:       int32(each.Index),
			OwnerId:     each.OwnerID,
			Visibility:  each.Visibility,
		}
		list = append(list, &r)
	}
	return &pb.ListProblemReply{
		Problems: list,
	}, nil
}
