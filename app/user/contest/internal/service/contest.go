package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "sastoj/api/sastoj/user/contest/service/v1"
)

func (s *UserContestService) ListContest(ctx context.Context, _ *pb.ListContestRequest) (*pb.ListContestReply, error) {
	// TODO: Get the userID from context
	userId := 1
	rv, err := s.contestUc.ListContest(ctx, int64(userId))
	if err != nil {
		return nil, err
	}
	reply := &pb.ListContestReply{}
	for _, p := range rv {
		reply.Contests = append(reply.Contests, &pb.Contest{
			Id:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			State:       int32(p.Status),
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

func (s *UserContestService) GetRanking(ctx context.Context, req *pb.GetRankingRequest) (*pb.GetRankingReply, error) {
	return &pb.GetRankingReply{}, nil
}
