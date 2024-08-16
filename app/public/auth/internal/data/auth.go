package data

import (
	"context"
	"sastoj/app/public/auth/internal/biz"
	"sastoj/ent/contest"
	"sastoj/ent/group"
	"sastoj/ent/user"

	"github.com/go-kratos/kratos/v2/log"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func (a *authRepo) GetContests(ctx context.Context, groupID int64) ([]int64, error) {
	contests, err := a.data.db.Contest.Query().Select(contest.FieldID).Where(contest.HasContestantsWith(group.IDEQ(groupID))).All(ctx)
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
	po, err := a.data.db.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		a.log.Errorf("FindUserByName failed: %v", err)
		return nil, err
	}
	g, err := po.QueryGroups().Only(ctx)
	return &biz.User{
		ID:        po.ID,
		Username:  po.Username,
		Password:  po.Password,
		Salt:      po.Salt,
		Status:    po.Status,
		GroupID:   po.GroupID,
		GroupName: g.GroupName,
	}, nil
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
