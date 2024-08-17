package service

import (
	"context"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/pkg/middleware/auth"
	"sastoj/pkg/util"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ContestService) Submit(ctx context.Context, req *pb.SubmitRequest) (*pb.SubmitReply, error) {
	caseVer, err := s.getProblemCaseVer(ctx, req.ProblemId)
	claim := ctx.Value("userInfo").(*auth.Claims)
	userID := claim.UserID
	id := uuid.NewString()
	if err != nil {
		return nil, err
	}
	err = s.submitUc.CreateSubmission(ctx, &biz.Submission{
		ID:          id,
		UserID:      userID,
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
		Uuid: id,
	}, nil
}

func (s *ContestService) SelfTest(ctx context.Context, req *pb.SelfTestRequest) (*pb.SelfTestReply, error) {
	selfTestID := uuid.NewString()
	claim := ctx.Value("userInfo").(*auth.Claims)
	userID := claim.UserID
	err := s.submitUc.CreateSelfTest(ctx, &biz.SelfTest{
		ID:       selfTestID,
		UserID:   userID,
		Code:     util.Crlf2lf(req.Code),
		Language: req.Language,
		Input:    util.Crlf2lf(req.Input),
	})
	if err != nil {
		return nil, err
	}
	return &pb.SelfTestReply{
		Uuid: selfTestID,
	}, nil
}

func (s *ContestService) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	submission, err := s.submitUc.GetSubmission(ctx, req.SubmissionId, req.GetContestId())
	if err != nil {
		return nil, err
	}
	return &pb.GetSubmissionReply{
		Id:        submission.ID,
		Code:      submission.Code,
		Language:  submission.Language,
		Point:     int32(submission.Point),
		State:     int32(submission.Status),
		CreatedAt: timestamppb.New(submission.CreateTime),
		TotalTime: submission.TotalTime,
		MaxMemory: submission.MaxMemory,
	}, nil
}

func (s *ContestService) GetSubmissions(ctx context.Context, req *pb.GetSubmissionsRequest) (*pb.GetSubmissionsReply, error) {
	submissions, err := s.submitUc.GetSubmissions(ctx, req.GetContestId(), req.GetProblemId())
	if err != nil {
		return nil, err
	}
	var pbSubmissions []*pb.GetSubmissionsReply_Submission
	for _, s := range submissions {
		pbSubmissions = append(pbSubmissions, &pb.GetSubmissionsReply_Submission{
			Id:        util.ParseInt64(s.ID),
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

func (s *ContestService) GetSelfTest(ctx context.Context, req *pb.GetSelfTestRequest) (*pb.GetSelfTestReply, error) {
	selfTest, err := s.submitUc.GetSelfTest(ctx, req.GetSelfTestId())
	if err != nil {
		return nil, err
	}
	return &pb.GetSelfTestReply{
		Id:       selfTest.ID,
		Code:     selfTest.Code,
		Language: selfTest.Language,
		Input:    selfTest.Input,
		Output:   selfTest.Output,
		Time:     selfTest.Time,
		Memory:   selfTest.Memory,
	}, nil
}

func (s *ContestService) GetCases(ctx context.Context, req *pb.GetCasesRequest) (*pb.GetCasesReply, error) {
	cases, err := s.submitUc.GetCases(ctx, req.GetSubmissionId(), req.GetContestId())
	if err != nil {
		return nil, err
	}
	var pbCases []*pb.GetCasesReply_Case
	for _, c := range cases {
		pbCases = append(pbCases, &pb.GetCasesReply_Case{
			Index: c.Index,
			State: c.State,
		})
	}
	return &pb.GetCasesReply{
		Cases: pbCases,
	}, nil
}
