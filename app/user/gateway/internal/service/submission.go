package service

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sastoj/app/user/gateway/internal/biz"
	"strconv"
	"time"

	pb "sastoj/api/sastoj/user/gateway/service/v1"
)

func (s *GatewayService) Submit(ctx context.Context, req *pb.SubmitRequest) (*pb.SubmitReply, error) {
	// TODO: Get the userID from context
	rv, err := s.submissionUc.CreateSubmission(ctx, &biz.Submission{
		ID:         uuid.New().String(),
		UserID:     0,
		ProblemID:  req.ProblemId,
		Code:       req.Code,
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
		Code:     req.Code,
		Language: req.Language,
		Input:    req.Input,
	})
	if err != nil {
		return nil, err
	}
	return &pb.SelfTestReply{
		Uuid: id,
	}, nil
}
func (s *GatewayService) GetSubmissions(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	// TODO: Get the userID from context
	rv, err := s.submissionUc.GetSubmission(ctx, strconv.FormatInt(req.SubmissionId, 10), 0)
	if err != nil {
		return nil, err
	}
	return &pb.GetSubmissionReply{
		Uuid:      rv.ID,
		Code:      rv.Code,
		Language:  rv.Language,
		Point:     int32(rv.Point),
		State:     int32(rv.Status),
		CreatedAt: timestamppb.New(rv.CreateTime),
		TotalTime: rv.TotalTime,
		MaxMemory: rv.MaxMemory,
	}, nil
}
func (s *GatewayService) ListRanking(ctx context.Context, req *pb.ListRankingRequest) (*pb.ListRankingReply, error) {
	return &pb.ListRankingReply{}, nil
}
