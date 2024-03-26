package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sastoj/app/user/contest/internal/biz"
	"strconv"

	pb "sastoj/api/sastoj/user/contest/service/v1"
)

type ContestService struct {
	pb.UnimplementedContestServiceServer
	uc *biz.ContestUsecase
}

func NewContestService(usecase *biz.ContestUsecase) *ContestService {
	return &ContestService{
		uc: usecase,
	}
}

func (s *ContestService) ListContest(ctx context.Context, req *pb.ListContestRequest) (*pb.ListContestReply, error) {
	// TODO: Get the userID from context
	rv, err := s.uc.ListContest(ctx, 1)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListContestReply{}
	for _, p := range rv {
		reply.Contests = append(reply.Contests, &pb.Contest{
			Id:          strconv.FormatInt(p.ID, 10),
			Title:       p.Title,
			Description: p.Description,
			State:       int32(p.State),
			Type:        int32(p.Type),
			StartTime:   timestamppb.New(p.StartTime),
			EndTime:     timestamppb.New(p.EndTime),
			Language:    p.Language,
			ExtraTime:   int32(p.ExtraTime),
		})
	}
	return reply, nil
}
func (s *ContestService) JoinContest(ctx context.Context, req *pb.JoinContestRequest) (*pb.JoinContestReply, error) {
	return &pb.JoinContestReply{}, nil
}
func (s *ContestService) GetProblem(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemReply, error) {
	return &pb.GetProblemReply{}, nil
}
func (s *ContestService) SubmitProblem(ctx context.Context, req *pb.SubmitProblemRequest) (*pb.SubmitProblemReply, error) {
	return &pb.SubmitProblemReply{}, nil
}
func (s *ContestService) PretestProblem(ctx context.Context, req *pb.PretestProblemRequest) (*pb.PretestProblemReply, error) {
	return &pb.PretestProblemReply{}, nil
}
func (s *ContestService) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	return &pb.GetSubmissionReply{}, nil
}
func (s *ContestService) GetRanking(ctx context.Context, req *pb.GetRankingRequest) (*pb.GetRankingReply, error) {
	return &pb.GetRankingReply{}, nil
}
