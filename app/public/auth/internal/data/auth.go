package data

import (
	"context"
	"sastoj/ent/user"

	"sastoj/app/public/auth/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &AuthRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AuthRepo) FindByUserName(ctx context.Context, username string) (*biz.User, error) {
	user, err := r.data.db.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return nil, err
	}
	group, err := user.QueryGroups().Only(ctx)
	return &biz.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Salt:      user.Salt,
		Status:    user.Status,
		GroupID:   user.GroupID,
		GroupName: group.GroupName,
	}, nil
}
