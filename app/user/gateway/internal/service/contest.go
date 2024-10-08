package service

import (
	"context"
	pb "sastoj/api/sastoj/user/contest/service/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GatewayService) GetContests(ctx context.Context, _ *pb.GetContestsRequest) (*pb.GetContestsReply, error) {
	// TODO: get groupID from ctx
	groupID := int64(0)
	rv, err := s.contestUc.GetContests(ctx, groupID)
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
		})

	}
	return reply, nil
}
func (s *GatewayService) JoinContest(ctx context.Context, req *pb.JoinContestRequest) (*pb.JoinContestReply, error) {
	return &pb.JoinContestReply{}, nil
}

func (s *GatewayService) ListRanking(ctx context.Context, req *pb.ListRankingRequest) (*pb.ListRankingReply, error) {
	return &pb.ListRankingReply{}, nil
}
