package service

import (
	"context"
	"github.com/google/uuid"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/contest/internal/biz"
	"time"
)

func (s *UserContestService) SubmitProblem(ctx context.Context, req *pb.SubmitProblemRequest) (*pb.SubmitProblemReply, error) {
	caseVer, err := s.getProblemCaseVer(ctx, int(req.ProblemId))
	if err != nil {
		return nil, err
	}
	submit, err := s.submitUc.CreateSubmit(ctx, &biz.Submit{
		UserID:      1, // TODO: Get the userID from context
		ProblemID:   int(req.ProblemId),
		Code:        req.Code,
		State:       0,
		Point:       0,
		CreateTime:  time.Now(),
		TotalTime:   0,
		MaxMemory:   0,
		Language:    req.Language,
		CaseVersion: caseVer,
	})
	if err != nil {
		return nil, err
	}
	return &pb.SubmitProblemReply{
		SubmitId: int64(submit),
	}, nil
}

func (s *UserContestService) PretestProblem(ctx context.Context, req *pb.PretestProblemRequest) (*pb.PretestProblemReply, error) {
	pretestId := uuid.New().String()
	err := s.submitUc.PretestProblem(ctx, &biz.Pretest{
		ID:       pretestId,
		UserID:   1, // TODO: Get the userID from context
		Code:     req.Code,
		Language: req.Language,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PretestProblemReply{
		PretestId: pretestId,
	}, nil
}
func (s *UserContestService) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	return &pb.GetSubmissionReply{}, nil
}
