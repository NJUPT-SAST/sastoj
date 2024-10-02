package data

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
	"sastoj/ent/user"
	"sastoj/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
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

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	res, err := r.data.db.User.Create().
		SetUsername(u.Username).
		SetPassword(u.Password).
		SetSalt(u.Salt).
		AddGroupIDs(u.GroupIDs...).
		Save(ctx)
	if err != nil {
		log.Debug("err: ", err)
		return nil, err
	}
	u.ID = res.ID
	return u, nil
}

func (r *userRepo) Update(ctx context.Context, u *biz.User) (*int64, error) {
	entState, err := util.UserStateToEnt(u.State)
	if err != nil {
		return nil, err
	}
	res, err := r.data.db.User.Update().
		SetUsername(u.Username).
		ClearGroups().
		AddGroupIDs(u.GroupIDs...).
		SetState(entState).
		Where(user.IDEQ(u.ID)).
		Save(ctx)
	if err != nil {
		log.Debug("err: ", err)
		return nil, err
	}
	res64 := int64(res)
	return &res64, nil
}

func (r *userRepo) FindByID(ctx context.Context, id int64) (*biz.User, error) {
	res, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		log.Debug("err: ", err)
		return nil, err
	}
	return &biz.User{
		ID:       res.ID,
		Username: res.Username,
		GroupIDs: res.QueryGroups().IDsX(ctx),
		State:    util.UserStateToInt(res.State),
	}, nil
}

func (r *userRepo) ListPages(ctx context.Context, current int64, size int64) ([]*biz.User, error) {
	res, err := r.data.db.User.Query().Offset(int((current - 1) * size)).Limit(int(size)).All(ctx)
	if err != nil {
		log.Debug("err: ", err)
		return nil, err
	}
	rv := make([]*biz.User, 0)
	for _, u := range res {
		rv = append(rv, &biz.User{
			ID:       u.ID,
			Username: u.Username,
			GroupIDs: u.QueryGroups().IDsX(ctx),
			State:    util.UserStateToInt(u.State),
		})
	}
	return rv, nil
}
func (r *userRepo) BatchSave(ctx context.Context, users []*biz.UserCreate) ([]string, error) {
	createdUsers, err := r.data.db.User.MapCreateBulk(users, func(c *ent.UserCreate, i int) {
		c.SetUsername(users[i].Username).
			SetSalt(users[i].Salt).
			SetPassword(users[i].Password).
			AddGroupIDs(users[i].GroupIDs...)
	}).Save(ctx)
	if err != nil {
		log.Debug("err: ", err)
	}
	// 返回usernames是为了防止有些账户没有创建成功
	usernames := make([]string, 0)
	for _, u := range createdUsers {
		usernames = append(usernames, u.Username)
	}
	return usernames, err
}
