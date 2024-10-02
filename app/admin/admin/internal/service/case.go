package service

import (
	"context"
	"fmt"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"strconv"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func (s *AdminService) DeleteCasesByProblemId(ctx context.Context, req *pb.DeleteCasesByProblemIdRequest) (*pb.DeleteCasesByProblemIdReply, error) {
	err := s.casec.DeleteCasesByProblemID(ctx, req.ProblemId)
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
	return ctx.String(200, fmt.Sprintf("Upload %s cases for problem %s successfully", casesType, problemId))
}
