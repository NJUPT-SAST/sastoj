package service

import (
	"context"
	"sastoj/app/admin/case/internal/biz"

	pb "sastoj/api/sastoj/admin/case/service/v1"
)

type CaseService struct {
	pb.UnimplementedCaseServiceServer
	uc *biz.CaseUsecase
}

func NewCaseService(usecase *biz.CaseUsecase) *CaseService {
	return &CaseService{
		uc: usecase,
	}
}

func (s *CaseService) CreateCases(ctx context.Context, req *pb.CreateCasesRequest) (*pb.CreateCasesReply, error) {
	var b []*biz.Case
	for _, i := range req.Requests {
		b = append(b, &biz.Case{
			ProblemId: req.GetProblemId(),
			Point:     i.Point,
			Index:     i.Index,
			IsAuto:    i.IsAuto,
		})
	}
	ids, err := s.uc.Save(ctx, req.GetProblemId(), b)
	return &pb.CreateCasesReply{
		Ids: ids,
	}, err
}
func (s *CaseService) UpdateCase(ctx context.Context, req *pb.UpdateCaseRequest) (*pb.UpdateCaseReply, error) {
	c := biz.Case{
		Id:        req.Id,
		ProblemId: req.ProblemId,
		Point:     req.Point,
		Index:     req.Index,
		IsAuto:    req.IsAuto,
	}
	err := s.uc.UpdateCase(ctx, &c)
	return &pb.UpdateCaseReply{}, err
}
func (s *CaseService) DeleteCasesByCaseIds(ctx context.Context, req *pb.DeleteCaseByCaseIdsRequest) (*pb.DeleteCaseByCaseIdsReply, error) {
	err := s.uc.DeleteCasesByCaseIds(ctx, req.Ids)
	return &pb.DeleteCaseByCaseIdsReply{}, err
}
func (s *CaseService) DeleteCasesByProblemId(ctx context.Context, req *pb.DeleteCasesByProblemIdRequest) (*pb.DeleteCasesByProblemIdReply, error) {
	err := s.uc.DeleteCasesByProblemId(ctx, req.ProblemId)
	return &pb.DeleteCasesByProblemIdReply{}, err
}
func (s *CaseService) GetCases(ctx context.Context, req *pb.GetCasesRequest) (*pb.GetCasesReply, error) {
	ps, err := s.uc.GetCasesByProblemId(ctx, req.GetProblemId())
	reply := &pb.GetCasesReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Case{
			Id:     p.Id,
			Point:  p.Point,
			Index:  p.Index,
			IsAuto: p.IsAuto,
		})
	}
	return reply, err
}
func (s *CaseService) UploadCases(conn pb.CaseService_UploadCasesServer) error {

	return nil
}
