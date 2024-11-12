package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent/group"
	"sastoj/pkg/util"
)

type groupRepo struct {
	data *Data
	log  *log.Helper
}

// NewGroupRepo .
func NewGroupRepo(data *Data, logger log.Logger) biz.GroupRepo {
	return &groupRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *groupRepo) Save(ctx context.Context, g *biz.Group) (*biz.Group, error) {
	po, err := r.data.db.Group.
		Create().
		SetGroupName(g.Name).
		AddManageIDs(g.ManageId...).
		AddContestIDs(g.ContestsId...).
		AddProblemIDs(g.ProblemsId...).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	r.log.WithContext(ctx).Infof("Save Group: %v", g.Name)
	return &biz.Group{Id: po.ID, Name: po.GroupName}, nil
}

func (r *groupRepo) Update(ctx context.Context, g *biz.Group) (*int64, error) {
	res, err := r.data.db.Group.Update().
		SetGroupName(g.Name).
		ClearManage().AddManageIDs(g.ManageId...).
		ClearContests().AddContestIDs(g.ContestsId...).
		ClearProblems().AddProblemIDs(g.ProblemsId...).
		Where(group.IDEQ(g.Id)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	res64 := int64(res)
	return &res64, nil
}

func (r *groupRepo) FindByID(ctx context.Context, id int64) (*biz.Group, error) {
	group, err := r.data.db.Group.Query().Where(group.IDEQ(id)).WithContests().WithProblems().WithManage().Only(ctx)
	if err != nil {
		return nil, err
	}
	log.Info(group.Edges.Problems)
	var manage []biz.Contest
	var contests []biz.Contest
	var problems []biz.Problem
	for _, m := range group.Edges.Manage {
		manage = append(manage, biz.Contest{
			Id:          m.ID,
			Title:       m.Title,
			Description: m.Description,
			Status:      util.ContestStateToInt(m.State, m.StartTime, m.EndTime),
			Type:        int32(m.Type),
			StartTime:   m.StartTime,
			EndTime:     m.EndTime,
			Language:    m.Language,
			ExtraTime:   int32(m.ExtraTime),
			CreateTime:  m.CreateTime,
		})
	}
	for _, c := range group.Edges.Contests {
		contests = append(contests, biz.Contest{
			Id:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			Status:      util.ContestStateToInt(c.State, c.StartTime, c.EndTime),
			Type:        int32(c.Type),
			StartTime:   c.StartTime,
			EndTime:     c.EndTime,
			Language:    c.Language,
			ExtraTime:   int32(c.ExtraTime),
			CreateTime:  c.CreateTime,
		})
	}
	for _, p := range group.Edges.Problems {
		problems = append(problems, biz.Problem{
			ID:          p.ID,
			TypeID:      p.ProblemTypeID,
			Title:       p.Title,
			Content:     p.Content,
			Point:       int32(p.Score),
			ContestID:   p.ContestID,
			CaseVersion: int32(p.CaseVersion),
			Index:       int32(p.Index),
		})
	}
	return &biz.Group{
		Id:       group.ID,
		Name:     group.GroupName,
		Manage:   manage,
		Contests: contests,
		Problems: problems,
	}, nil
}

func (r *groupRepo) ListByPage(ctx context.Context, current int64, size int64) ([]*biz.Group, error) {
	res, err := r.data.db.Group.Query().Limit(int(size)).Offset(int((current - 1) * size)).All(ctx)
	if err != nil {
		return nil, err
	}
	var groups []*biz.Group
	for _, g := range res {
		groups = append(groups, &biz.Group{
			Id:   g.ID,
			Name: g.GroupName,
		})
	}
	return groups, nil
}

func (r *groupRepo) DeleteById(ctx context.Context, id int64) error {
	err := r.data.db.Group.DeleteOneID(id).Exec(ctx)
	return err
}
