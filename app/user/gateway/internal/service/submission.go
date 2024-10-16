package service

import (
	"context"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/gateway/internal/biz"
	"sastoj/pkg/util"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GatewayService) Submit(ctx context.Context, req *pb.SubmitRequest) (*pb.SubmitReply, error) {
	// TODO: Get the userID from context
	rv, err := s.submissionUc.CreateSubmission(ctx, &biz.Submission{
		ID:         uuid.New().String(),
		UserID:     0,
		ProblemID:  req.ProblemId,
		Code:       util.Crlf2lf(req.Code),
		Status:     0,
		Point:      0,
		CreateTime: time.Now(),
		TotalTime:  0,
		MaxMemory:  0,
		Language:   req.Language,
	})
	if err != nil {
		return nil, err
	}
	return &pb.SubmitReply{
		Uuid: rv,
	}, nil
}

func (s *GatewayService) SelfTest(ctx context.Context, req *pb.SelfTestRequest) (*pb.SelfTestReply, error) {
	// TODO: Get the userID from context
	id := uuid.New().String()
	err := s.submissionUc.CreateSelfTest(ctx, &biz.SelfTest{
		ID:       id,
		UserID:   0,
		Code:     util.Crlf2lf(req.Code),
		Language: req.Language,
		Input:    util.Crlf2lf(req.Input),
	})
	if err != nil {
		return nil, err
	}
	return &pb.SelfTestReply{
		Uuid: id,
	}, nil
}

func (s *GatewayService) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	rv, err := s.submissionUc.GetSubmission(ctx, req.GetSubmissionId(), req.GetContestId())
	if err != nil {
		return nil, err
	}
	return &pb.GetSubmissionReply{
		Id:        rv.ID,
		Code:      rv.Code,
		Language:  rv.Language,
		Point:     int32(rv.Point),
		State:     int32(rv.Status),
		CreatedAt: timestamppb.New(rv.CreateTime),
		TotalTime: uint64(rv.TotalTime),
		MaxMemory: uint64(rv.MaxMemory),
	}, nil
}

func (s *GatewayService) GetSubmissions(ctx context.Context, req *pb.GetSubmissionsRequest) (*pb.GetSubmissionsReply, error) {
	submissions, err := s.submissionUc.GetSubmissions(ctx, req.ContestId, req.ProblemId)
	if err != nil {
		return nil, err
	}
	res := &pb.GetSubmissionsReply{}
	for _, v := range submissions {
		res.Submissions = append(res.Submissions, &pb.GetSubmissionsReply_Submission{
			Id:        util.ParseInt64(v.ID),
			Language:  v.Language,
			Point:     int32(v.Point),
			Status:    int32(v.Status),
			CreatedAt: timestamppb.New(v.CreateTime),
		})
	}
	return res, nil
}

func (s *GatewayService) GetCases(ctx context.Context, req *pb.GetCasesRequest) (*pb.GetCasesReply, error) {
	cases, err := s.submissionUc.GetCases(ctx, req.ContestId, req.SubmissionId)
	if err != nil {
		return nil, err
	}
	res := &pb.GetCasesReply{}
	for _, v := range cases {
		res.Cases = append(res.Cases, &pb.GetCasesReply_Case{
			Index: v.Index,
			State: v.State,
		})
	}
	return res, nil
}

func (s *GatewayService) UpdateSubmission(ctx context.Context, req *pb.UpdateSubmissionRequest) (*pb.UpdateSubmissionReply, error) {
	err := s.submissionUc.UpdateSubmission(ctx, &biz.Submission{
		ID:        req.SubmissionId,
		Status:    int16(req.Status),
		Point:     int16(*req.Point),
		TotalTime: *req.TotalTime,
		MaxMemory: *req.MaxMemory,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateSubmissionReply{
		Success: true,
	}, nil
}

func (s *GatewayService) UpdateSelfTest(ctx context.Context, req *pb.UpdateSelfTestRequest) (*pb.UpdateSelfTestReply, error) {
	err := s.submissionUc.UpdateSelfTest(ctx, &biz.SelfTest{
		ID:     req.SelfTestId,
		Output: req.Output,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateSelfTestReply{
		Success: true,
	}, nil
}
