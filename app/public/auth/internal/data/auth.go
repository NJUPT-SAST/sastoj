package data

import (
	"context"
	"sastoj/app/public/auth/internal/biz"
	"sastoj/ent"
	"sastoj/ent/contest"
	"sastoj/ent/group"
	"sastoj/ent/user"
	"sastoj/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func (a *authRepo) GetContests(ctx context.Context, groupIds []int64) ([]int64, error) {
	contests, err := a.data.db.Contest.Query().Select(contest.FieldID).Where(contest.HasContestantsWith(group.IDIn(groupIds...))).All(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, c := range contests {
		ids = append(ids, c.ID)
	}
	return ids, nil
}

func (a *authRepo) FindUserByName(ctx context.Context, username string) (*biz.User, error) {
	po, err := a.data.db.User.Query().
		Where(
			user.UsernameEQ(username)).
		Select(user.FieldID, user.FieldUsername, user.FieldPassword, user.FieldSalt, user.FieldState).
		WithGroups(func(q *ent.GroupQuery) {
			q.Select(group.FieldID, group.FieldGroupName)
		}).
		Only(ctx)
	if err != nil {
		a.log.Errorf("FindUserByName failed: %v", err)
		return nil, err
	}
	groupIDs := make([]int64, 0)
	for _, g := range po.Edges.Groups {
		groupIDs = append(groupIDs, g.ID)
	}
	return &biz.User{
		ID:        po.ID,
		Username:  po.Username,
		Password:  po.Password,
		Salt:      po.Salt,
		Status:    util.UserStateToInt(po.State),
		GroupIds:  groupIDs,
		GroupName: po.Edges.Groups[0].GroupName,
	}, nil
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
