package service

import (
	"context"
	"errors"
	v1 "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
	"sastoj/pkg/util"
)

func (s *AdminService) CreateContest(ctx context.Context, req *v1.CreateContestRequest) (*v1.CreateContestReply, error) {
	rv, err := s.ctsc.CreateContest(ctx, &biz.Contest{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Type:        req.Type,
		StartTime:   req.StartTime.AsTime(),
		EndTime:     req.EndTime.AsTime(),
		Language:    req.Language,
		ExtraTime:   req.ExtraTime,
		CreateTime:  req.CreateTime.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateContestReply{
		Id: rv.Id,
	}, nil
}
func (s *AdminService) UpdateContest(ctx context.Context, req *v1.UpdateContestRequest) (*v1.UpdateContestReply, error) {
	err := s.ctsc.UpdateContest(ctx, &biz.Contest{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Type:        req.Type,
		StartTime:   req.StartTime.AsTime(),
		EndTime:     req.EndTime.AsTime(),
		Language:    req.Language,
		ExtraTime:   req.ExtraTime,
		CreateTime:  req.CreateTime.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateContestReply{
		Success: true,
	}, nil
}
func (s *AdminService) DeleteContest(ctx context.Context, req *v1.DeleteContestRequest) (*v1.DeleteContestReply, error) {
	err := s.ctsc.DeleteContest(ctx, req.Id)
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			return nil, v1.ErrorContestNotFound("contest with specified ID not found")
		}
		return nil, err
	}
	return &v1.DeleteContestReply{
		Success: true,
	}, nil
}
func (s *AdminService) GetContest(ctx context.Context, req *v1.GetContestRequest) (*v1.GetContestReply, error) {
	rv, err := s.ctsc.FindContest(ctx, req.Id)
	if err != nil {

		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			return nil, v1.ErrorContestNotFound("contest with specified ID not found")
		}
		return nil, err
	}
	return &v1.GetContestReply{
		Id:          rv.Id,
		Title:       rv.Title,
		Description: rv.Description,
		Status:      rv.Status,
		Type:        rv.Type,
		StartTime:   util.ConvertTimeToTimeStamp(rv.StartTime),
		EndTime:     util.ConvertTimeToTimeStamp(rv.EndTime),
		Language:    rv.Language,
		ExtraTime:   rv.ExtraTime,
		CreateTime:  util.ConvertTimeToTimeStamp(rv.CreateTime),
	}, nil
}
func (s *AdminService) ListContest(ctx context.Context, req *v1.ListContestRequest) (*v1.ListContestReply, error) {
	rv, total, err := s.ctsc.ListContest(ctx, req.Current, req.Size)
	if err != nil {
		return nil, err
	}
	var list []*v1.ListContestReply_Contest
	for _, v := range rv {
		list = append(list, &v1.ListContestReply_Contest{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Status:      v.Status,
			Type:        v.Type,
			StartTime:   util.ConvertTimeToTimeStamp(v.StartTime),
			EndTime:     util.ConvertTimeToTimeStamp(v.EndTime),
			Language:    v.Language,
			ExtraTime:   v.ExtraTime,
			CreateTime:  util.ConvertTimeToTimeStamp(v.CreateTime),
		})
	}
	return &v1.ListContestReply{
		Contests: list,
		Total:    total,
	}, nil
}
func (s *AdminService) AddContestants(ctx context.Context, req *v1.AddContestantsRequest) (*v1.AddContestantsReply, error) {
	if req.Role != 0 && req.Role != 1 {
		return nil, v1.ErrorInvalidArgument("role must be 0 or 1")
	}
	err := s.ctsc.AddContestants(ctx, req.ContestId, req.GroupId, req.Role)
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			return nil, v1.ErrorContestNotFound("contest with specified ID not found")
		}
		return nil, v1.ErrorGroupNotFound("group with specified ID not found")
	}
	return &v1.AddContestantsReply{
		Success: true,
	}, nil
}
func (s *AdminService) ManualRanking(ctx context.Context, req *v1.ManualRankingRequest) (*v1.ManualRankingReply, error) {
	contest, err := s.ctsc.FindContest(ctx, req.ContestId)
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			return nil, v1.ErrorContestNotFound("contest with specified ID not found")
		}
		return nil, err
	}
	rank, err := s.rc.Update(ctx, contest)
	if err != nil {
		return nil, err
	}
	err = s.rc.Save(ctx, rank)
	if err != nil {
		return nil, err
	}
	err = s.rc.SaveCache(ctx, contest, rank)
	return &v1.ManualRankingReply{Success: true}, nil
}

func (s *AdminService) GetRanking(ctx context.Context, req *v1.GetRankingRequest) (*v1.GetRankingReply, error) {
	contest, err := s.ctsc.FindContest(ctx, req.ContestId)
	if err != nil {
		var entErr *ent.NotFoundError
		if errors.As(err, &entErr) {
			return nil, v1.ErrorContestNotFound("contest with specified ID not found")
		}
		return nil, err
	}
	rv, err := s.rc.Find(ctx, contest)
	if err != nil {
		return nil, err
	}
	list := make([]*v1.GetRankingReply_UserResult, 0)
	for _, v := range rv.UserRank {
		problems := make([]*v1.GetRankingReply_UserResult_ProblemResult, 0)
		for _, p := range v.Problems {
			problems = append(problems, &v1.GetRankingReply_UserResult_ProblemResult{
				ProblemId:         p.ProblemId,
				State:             p.State,
				Point:             p.Point,
				TriedTimes:        p.TriedCount,
				ScoreAchievedTime: util.ConvertTimeToTimeStamp(p.SubmitTime),
			})
		}
		list = append(list, &v1.GetRankingReply_UserResult{
			Username:   v.UserName,
			TotalScore: v.TotalPoint,
			Rank:       v.Rank,
			Penalty:    v.Penalty,
			Problems:   problems,
		})
	}
	return &v1.GetRankingReply{
		Users: list,
	}, nil
}
