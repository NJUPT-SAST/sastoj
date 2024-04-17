package service

import (
	"context"
	pb "sastoj/api/sastoj/admin/problem/service/v1"
	"sastoj/app/admin/problem/internal/biz"
)

type ProblemService struct {
	pb.UnimplementedProblemServer
	pu *biz.ProblemUsecase
}

func NewProblemService(problem *biz.ProblemUsecase) *ProblemService {
	return &ProblemService{pu: problem}
}

func (s *ProblemService) CreateProblem(ctx context.Context, request *pb.CreateProblemRequest) (*pb.CreateProblemReply, error) {
	rv, err := s.pu.CreateProblem(ctx, &biz.Problem{
		Title:       request.Title,
		Content:     request.Content,
		Point:       request.Point,
		ContestId:   request.ContestId,
		CaseVersion: request.CaseVersion,
		Index:       request.Index,
		Config:      request.Config,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateProblemReply{
		Id: rv.Id,
	}, nil
}

func (s *ProblemService) UpdateProblem(ctx context.Context, request *pb.UpdateProblemRequest) (*pb.UpdateProblemReply, error) {
	rv, err := s.pu.UpdateProblem(ctx, &biz.Problem{
		Id:          request.Id,
		Title:       request.Title,
		Content:     request.Content,
		Point:       request.Point,
		ContestId:   request.ContestId,
		CaseVersion: request.CaseVersion,
		Index:       request.Index,
		Config:      request.Config,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProblemReply{
		Success: rv,
	}, nil
}

func (s *ProblemService) DeleteProblem(ctx context.Context, request *pb.DeleteProblemRequest) (*pb.DeleteProblemReply, error) {
	rv, err := s.pu.DeleteProblem(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProblemReply{
		Success: rv,
	}, nil
}

func (s *ProblemService) GetProblem(ctx context.Context, request *pb.GetProblemRequest) (*pb.GetProblemReply, error) {
	rv, err := s.pu.FindProblem(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProblemReply{
		Id: rv.Id,
	}, nil
}

func (s *ProblemService) ListProblem(ctx context.Context, request *pb.ListProblemRequest) (*pb.ListProblemReply, error) {
	rv, err := s.pu.ListProblem(ctx, request.Currency, request.Size)
	if err != nil {
		return nil, err
	}
	var list []*pb.ListProblemReply_Problem
	for _, each := range rv {
		list = append(list, &pb.ListProblemReply_Problem{
			Id:          each.Id,
			Title:       each.Title,
			Content:     each.Content,
			Point:       each.Point,
			ContestId:   each.ContestId,
			CaseVersion: each.CaseVersion,
			Index:       each.Index,
			Config:      each.Config,
		})
	}
	return &pb.ListProblemReply{
		Problems: list,
	}, nil
}
