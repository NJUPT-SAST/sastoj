package service

import (
	"context"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/pkg/util"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ContestService) GetContests(ctx context.Context, _ *pb.GetContestsRequest) (*pb.GetContestsReply, error) {
	userID := util.GetUserInfoFromCtx(ctx).UserId
	rv, err := s.contestUc.ListContest(ctx, userID)
	if err != nil {
		return nil, err
	}
	reply := &pb.GetContestsReply{}
	for _, p := range rv {
		reply.Contests = append(reply.Contests, &pb.GetContestsReply_Contest{
			Id:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			State:       int32(p.Status),
			Type:        int32(p.Type),
			StartTime:   timestamppb.New(p.StartTime),
			EndTime:     timestamppb.New(p.EndTime),
			Language:    p.Language,
			ExtraTime:   int32(p.ExtraTime),
			Contestants: p.Groups,
		})
	}
	return reply, nil
}
func (s *ContestService) JoinContest(ctx context.Context, req *pb.JoinContestRequest) (*pb.JoinContestReply, error) {
	userID := util.GetUserInfoFromCtx(ctx).UserId
	err := s.contestUc.JoinContest(ctx, userID, req.GetContestId(), req.Body.IsJoin)
	if err != nil {
		return nil, err
	}
	return &pb.JoinContestReply{
		IsJoin: req.Body.IsJoin,
	}, nil
}

func (s *ContestService) ListRanking(ctx context.Context, req *pb.ListRankingRequest) (*pb.ListRankingReply, error) {
	return &pb.ListRankingReply{}, nil
}
