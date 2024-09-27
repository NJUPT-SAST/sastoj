package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"strconv"
)

func (s *AdminService) DeleteCasesByProblemId(ctx context.Context, req *pb.DeleteCasesByProblemIdRequest) (*pb.DeleteCasesByProblemIdReply, error) {
	err := s.casec.DeleteCasesByProblemId(ctx, req.ProblemId)
	return &pb.DeleteCasesByProblemIdReply{}, err
}

func (s *AdminService) UploadCases(ctx http.Context) error {
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
	err = s.casec.UploadCases(ctx, int64CaseId, file, filename, casesType)
	if err != nil {
		return err
	}
	return nil
}
