package service

import (
	"context"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/pkg/util"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserContestService) Submit(ctx context.Context, req *pb.SubmitRequest) (*pb.SubmitReply, error) {
	caseVer, err := s.getProblemCaseVer(ctx, req.ProblemId)
	if err != nil {
		return nil, err
	}
	submit, err := s.submitUc.CreateSubmit(ctx, &biz.Submit{
		UserID:      1, // TODO: Get the userID from context
		ProblemID:   req.ProblemId,
		Code:        util.Crlf2lf(req.Code),
		Status:      0,
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
	return &pb.SubmitReply{
		SubmitId: submit,
	}, nil
}

func (s *UserContestService) SelfTest(ctx context.Context, req *pb.SelfTestRequest) (*pb.SelfTestReply, error) {
	pretestId := uuid.New().String()
	err := s.submitUc.PretestProblem(ctx, &biz.Pretest{
		ID:       pretestId,
		UserID:   1, // TODO: Get the userID from context
		Code:     util.Crlf2lf(req.Code),
		Language: req.Language,
		Input:    util.Crlf2lf(req.Input),
	})
	if err != nil {
		return nil, err
	}
	return &pb.SelfTestReply{
		PretestId: pretestId,
	}, nil
}

func (s *UserContestService) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	submission, err := s.submitUc.GetSubmission(ctx, req.SubmissionId, 1) // TODO: Get the userID from context
	if err != nil {
		return nil, err
	}
	return &pb.GetSubmissionReply{
		Id:        submission.ID,
		Code:      submission.Code,
		Language:  submission.Language,
		Point:     int32(submission.Point),
		Status:    int32(submission.Status),
		CreatedAt: timestamppb.New(submission.CreateTime),
		TotalTime: submission.TotalTime,
		MaxMemory: submission.MaxMemory,
	}, nil
}

func (s *UserContestService) GetSubmissions(ctx context.Context, req *pb.GetSubmissionsRequest) (*pb.GetSubmissionsReply, error) {
	// TODO: Get the userID from context
	submissions, err := s.submitUc.GetSubmissions(ctx, 1, req.GetProblemId())
	if err != nil {
		return nil, err
	}
	var pbSubmissions []*pb.GetSubmissionsReply_Submission
	for _, s := range submissions {
		pbSubmissions = append(pbSubmissions, &pb.GetSubmissionsReply_Submission{
			Id:        s.ID,
			Language:  s.Language,
			Point:     int32(s.Point),
			Status:    int32(s.Status),
			CreatedAt: timestamppb.New(s.CreateTime),
		})
	}
	return &pb.GetSubmissionsReply{
		Submissions: pbSubmissions,
	}, nil
}

func (s *UserContestService) GetCases(ctx context.Context, req *pb.GetCasesRequest) (*pb.GetCasesReply, error) {
	var userID int64 = 1 // TODO: Get the userID from context
	cases, err := s.submitUc.GetCases(req.GetSubmissionId(), userID)
	if err != nil {
		return nil, err
	}
	var pbCases []*pb.GetCasesReply_Case
	for _, c := range cases {
		pbCases = append(pbCases, &pb.GetCasesReply_Case{
			Index:  int32(c.Index),
			Status: int32(c.State),
		})
	}
	return &pb.GetCasesReply{
		Cases: pbCases,
	}, nil
}
