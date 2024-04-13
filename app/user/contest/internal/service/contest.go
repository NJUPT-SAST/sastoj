package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "sastoj/api/sastoj/user/contest/service/v1"
)

func (s *UserContestService) ListContest(ctx context.Context, req *pb.ListContestRequest) (*pb.ListContestReply, error) {
	// TODO: Get the userID from context
	rv, err := s.contestUc.ListContest(ctx, 1)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListContestReply{}
	for _, p := range rv {
		reply.Contests = append(reply.Contests, &pb.Contest{
			Id:          p.ID,
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
func (s *UserContestService) JoinContest(ctx context.Context, req *pb.JoinContestRequest) (*pb.JoinContestReply, error) {
	err := s.contestUc.JoinContest(ctx, 1, 1, req.JoinContestBody.IsJoin)
	if err != nil {
		return nil, err
	}
	return &pb.JoinContestReply{
		IsJoin: req.JoinContestBody.IsJoin,
	}, nil
}

func (s *UserContestService) SubmitProblem(ctx context.Context, req *pb.SubmitProblemRequest) (*pb.SubmitProblemReply, error) {
	return &pb.SubmitProblemReply{}, nil
}

func (s *UserContestService) PretestProblem(ctx context.Context, req *pb.PretestProblemRequest) (*pb.PretestProblemReply, error) {
	return &pb.PretestProblemReply{}, nil
}
func (s *UserContestService) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionReply, error) {
	return &pb.GetSubmissionReply{}, nil
}
func (s *UserContestService) GetRanking(ctx context.Context, req *pb.GetRankingRequest) (*pb.GetRankingReply, error) {
	return &pb.GetRankingReply{}, nil
}
