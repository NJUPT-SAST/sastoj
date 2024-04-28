package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "sastoj/api/sastoj/user/gateway/service/v1"
)

func (s *GatewayService) GetContests(ctx context.Context, req *pb.GetContestsRequest) (*pb.GetContestsReply, error) {
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
