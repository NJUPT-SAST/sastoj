package service

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
	"time"
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
			return nil, v1.ErrorContestNotFound("contest with specified Id not found")
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
			return nil, v1.ErrorContestNotFound("contest with specified Id not found")
		}
		return nil, err
	}
	return &v1.GetContestReply{
		Id:          rv.Id,
		Title:       rv.Title,
		Description: rv.Description,
		Status:      rv.Status,
		Type:        rv.Type,
		StartTime:   convertTimeToTimeStamp(rv.StartTime),
		EndTime:     convertTimeToTimeStamp(rv.EndTime),
		Language:    rv.Language,
		ExtraTime:   rv.ExtraTime,
		CreateTime:  convertTimeToTimeStamp(rv.CreateTime),
	}, nil
}
func (s *AdminService) ListContest(ctx context.Context, req *v1.ListContestRequest) (*v1.ListContestReply, error) {
	rv, err := s.ctsc.ListContest(ctx, req.Current, req.Size)
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
			StartTime:   convertTimeToTimeStamp(v.StartTime),
			EndTime:     convertTimeToTimeStamp(v.EndTime),
			Language:    v.Language,
			ExtraTime:   v.ExtraTime,
			CreateTime:  convertTimeToTimeStamp(v.CreateTime),
		})
	}
	return &v1.ListContestReply{
		Contests: list,
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
			return nil, v1.ErrorContestNotFound("contest with specified Id not found")
		}
		return nil, v1.ErrorGroupNotFound("group with specified Id not found")
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
			return nil, v1.ErrorContestNotFound("contest with specified Id not found")
		}
		return nil, err
	}
	rank, err := s.rc.FindRank(ctx, contest)
	if err != nil && rank != nil {
		return nil, err
	}
	rank, err = s.rc.RefreshRank(ctx, contest, rank)
	if err != nil && rank != nil {
		return nil, err
	}
	err = s.rc.SaveRank(ctx, contest, rank)
	if err != nil {
		return nil, err
	}
	return &v1.ManualRankingReply{Success: true}, nil
}

func convertTimeToTimeStamp(tm time.Time) *timestamppb.Timestamp {

	return &timestamppb.Timestamp{
		Seconds: tm.Unix(),
		Nanos:   int32(tm.Nanosecond()),
	}
}
