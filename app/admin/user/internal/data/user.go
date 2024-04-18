package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/user/internal/biz"
	"sastoj/ent/user"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	res, err := r.data.db.User.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetSalt(user.Salt).
		SetState(user.State).
		SetGroupID(user.GroupID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	user.ID = res.ID
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, u *biz.User) (*int, error) {
	res, err := r.data.db.User.Update().
		SetUsername(u.Username).
		Where(user.ID(int(u.ID))).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *userRepo) FindByID(ctx context.Context, id int) (*biz.User, error) {
	res, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       res.ID,
		Username: res.Username,
	}, nil
}

func (r *userRepo) ListPages(ctx context.Context, current int, size int) ([]*biz.User, error) {
	res, err := r.data.db.User.Query().Offset((current - 1) * size).Limit(size).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.User, 0)
	for _, u := range res {
		rv = append(rv, &biz.User{
			ID:       u.ID,
			Username: u.Username,
		})
	}
	return rv, nil
}
