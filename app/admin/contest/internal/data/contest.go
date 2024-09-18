package data

import (
	"context"
	"errors"
	v1 "sastoj/api/sastoj/admin/contest/service/v1"
	"sastoj/app/admin/contest/internal/biz"
	"sastoj/ent/contest"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ContestRepo struct {
	data *Data
	log  *log.Helper
}

func NewContestRepo(data *Data, logger log.Logger) biz.ContestRepo {
	return &ContestRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ContestRepo) Save(ctx context.Context, g *biz.Contest) (*biz.Contest, error) {
	entState, err := stateToEntState(g.Status)
	if err != nil {
		return nil, err
	}
	res, err := r.data.db.Contest.Create().
		SetTitle(g.Title).
		SetDescription(g.Description).
		SetState(entState).
		SetType(int16(g.Type)).
		SetStartTime(g.StartTime).
		SetEndTime(g.EndTime).
		SetLanguage(g.Language).
		SetExtraTime(int16(g.ExtraTime)).
		SetCreateTime(g.CreateTime).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	g.Id = res.ID
	return g, nil
}

func (r *ContestRepo) Update(ctx context.Context, g *biz.Contest) error {
	entState, err := stateToEntState(g.Status)
	if err != nil {
		return err
	}
	_, err = r.data.db.Contest.Update().
		SetTitle(g.Title).
		SetDescription(g.Description).
		SetState(entState).
		SetType(int16(g.Type)).
		SetStartTime(g.StartTime).
		SetEndTime(g.EndTime).
		SetLanguage(g.Language).
		SetExtraTime(int16(g.ExtraTime)).
		SetCreateTime(g.CreateTime).
		Where(contest.ID(g.Id)).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContestRepo) FindByID(ctx context.Context, id int64) (*biz.Contest, error) {
	po, err := r.data.db.Contest.Get(ctx, id)
	if err != nil {
		log.Debug("err : ", err)
		return nil, err
	}
	return &biz.Contest{
		Id:          po.ID,
		Title:       po.Title,
		Description: po.Description,
		Status:      stateToInt(po.State, po.StartTime, po.EndTime),
		Type:        int32(po.Type),
		StartTime:   po.StartTime,
		EndTime:     po.EndTime,
		Language:    po.Language,
		ExtraTime:   int32(po.ExtraTime),
		CreateTime:  po.CreateTime,
	}, nil
}
func (r *ContestRepo) Delete(ctx context.Context, id int64) error {
	err := r.data.db.Contest.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContestRepo) ListPages(ctx context.Context, current int64, size int64) ([]*biz.Contest, error) {
	res, err := r.data.db.Contest.Query().Offset(int((current - 1) * size)).Limit(int(size)).All(ctx)
	if err != nil {
		log.Debug(" error :", err)
		return nil, err
	}
	rv := make([]*biz.Contest, 0)
	for _, po := range res {
		rv = append(rv, &biz.Contest{
			Id:          po.ID,
			Title:       po.Title,
			Description: po.Description,
			Status:      stateToInt(po.State, po.StartTime, po.EndTime),
			Type:        int32(po.Type),
			StartTime:   po.StartTime,
			EndTime:     po.EndTime,
			Language:    po.Language,
			ExtraTime:   int32(po.ExtraTime),
			CreateTime:  po.CreateTime,
		})
	}
	return rv, nil
}
func (r *ContestRepo) AddContestants(ctx context.Context, contestId int64, groupId int64, role int32) error {
	switch role {
	case 0:
		_, err := r.data.db.Contest.UpdateOneID(contestId).AddContestantIDs(groupId).Save(ctx)
		if err != nil {
			log.Debug(" error :", err)
			return err
		}
		return nil
	case 1:
		_, err := r.data.db.Contest.UpdateOneID(contestId).AddManagerIDs(groupId).Save(ctx)
		if err != nil {
			log.Debug(" error :", err)
			return err
		}
		return nil
	default:
		return errors.New("role not exist")
	}

}

func (r *ContestRepo) GetRacingContests(ctx context.Context) ([]*biz.Contest, error) {
	cs, err := r.data.db.Contest.Query().Where(
		contest.StartTimeLTE(time.Now()),
		contest.EndTimeGTE(time.Now().Add(time.Hour)),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Contest, 0)
	for _, c := range cs {
		rv = append(rv, &biz.Contest{
			Title:       c.Title,
			Description: c.Description,
			Status:      stateToInt(c.State, c.StartTime, c.EndTime),
			Type:        int32(c.Type),
			StartTime:   c.StartTime,
			EndTime:     c.EndTime,
			Language:    c.Language,
			ExtraTime:   int32(c.ExtraTime),
			CreateTime:  c.CreateTime,
		})
	}
	return rv, nil
}

func stateToInt(state contest.State, start, end time.Time) int32 {
	switch state {
	case contest.StateNORMAL:
		now := time.Now()
		switch {
		case now.Before(start):
			return 0
		case now.After(end):
			return 2
		default:
			return 1
		}
	case contest.StateCANCELLED:
		return 3
	case contest.StateHIDDEN:
		return 4
	case contest.StateDELETED:
		return 5
	default:
		return -1
	}
}

func stateToEntState(state int32) (contest.State, error) {
	switch state {
	case 0, 1, 2:
		return contest.StateNORMAL, nil
	case 3:
		return contest.StateCANCELLED, nil
	case 4:
		return contest.StateHIDDEN, nil
	case 5:
		return contest.StateDELETED, nil
	default:
		return "", v1.ErrorContestInvalid("contest type invalid")
	}
}
