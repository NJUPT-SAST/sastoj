package service

import (
	"context"
	"errors"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/ent"
	"sastoj/pkg/util"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ContestService) GetContests(ctx context.Context, _ *pb.GetContestsRequest) (*pb.GetContestsReply, error) {
	userID := util.GetUserInfoFromCtx(ctx).UserId
	exist := s.contestUc.CheckBanned(ctx, userID)
	if exist {
		return nil, pb.ErrorUserBanned("user is banned")
	}
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
	exist := s.contestUc.CheckBanned(ctx, userID)
	if exist {
		return nil, pb.ErrorUserBanned("user is banned")
	}
	err := s.contestUc.JoinContest(ctx, userID, req.GetContestId(), req.Body.IsJoin)
	if err != nil {
		return nil, err
	}
	return &pb.JoinContestReply{
		IsJoin: req.Body.IsJoin,
	}, nil
}

func (s *ContestService) ListRanking(ctx context.Context, req *pb.ListRankingRequest) (*pb.ListRankingReply, error) {
	rv, err := s.rankUc.Find(ctx, req.ContestId)
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			return nil, errors.New("contest with specified Id not found")
		}
		return nil, err
	}
	list := make([]*pb.ListRankingReply_UserResult, 0)
	for _, v := range rv.UserRank {
		problems := make([]*pb.ListRankingReply_UserResult_ProblemResult, 0)
		for _, p := range v.Problems {
			problems = append(problems, &pb.ListRankingReply_UserResult_ProblemResult{
				ProblemId:         p.ProblemId,
				State:             p.State,
				Point:             p.Point,
				TriedTimes:        p.TriedCount,
				ScoreAchievedTime: util.ConvertTimeToTimeStamp(p.SubmitTime),
			})
		}
		list = append(list, &pb.ListRankingReply_UserResult{
			Username:   v.UserName,
			TotalScore: v.TotalPoint,
			Rank:       v.Rank,
			Penalty:    v.Penalty,
			Problems:   problems,
		})
	}
	return &pb.ListRankingReply{
		Users: list,
	}, nil
}
