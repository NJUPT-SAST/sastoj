package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "sastoj/api/sastoj/admin/case/service/v1"
	"sastoj/app/admin/case/internal/biz"
	"strconv"
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
func (s *CaseService) DeleteCasesByProblemId(ctx context.Context, req *pb.DeleteCasesByProblemIdRequest) (*pb.DeleteCasesByProblemIdReply, error) {
	err := s.uc.DeleteCasesByProblemId(ctx, req.ProblemId)
	return &pb.DeleteCasesByProblemIdReply{}, err
}

func (s *CaseService) UploadCases(ctx http.Context) error {
	req := ctx.Request()
	casesType := req.FormValue("type")
	problemId := req.FormValue("problemId")
	int64CaseId, err := strconv.ParseInt(problemId, 10, 64)
	if err != nil {
		return err
	}
	file, handler, err := req.FormFile("file")
	if err != nil {
		return err
	}
	filename := handler.Filename
	defer file.Close()
	err = s.uc.UploadCases(ctx, int64CaseId, file, filename, casesType)
	if err != nil {
		return err
	}
	return nil
}
