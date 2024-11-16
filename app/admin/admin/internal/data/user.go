package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent"
	"sastoj/ent/group"
	"sastoj/ent/user"
	"sastoj/pkg/util"
)

const userStatePrefix = "user:contest:userState:"

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
		AddGroupIDs(u.GroupIds...).
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
		AddGroupIDs(u.GroupIds...).
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

func (r *userRepo) DeleteCache(ctx context.Context, id int64) error {
	return r.data.redis.SRem(ctx, userStatePrefix, id).Err()
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
		GroupIds: res.QueryGroups().IDsX(ctx),
		State:    util.UserStateToInt(res.State),
	}, nil
}

func (r *userRepo) ListPages(ctx context.Context, current int64, size int64, groupIDs []int64, username string, state int16) ([]*biz.User, int64, error) {
	query := r.data.db.User.Query()
	if len(groupIDs) > 0 {
		query = query.Where(user.HasGroupsWith(group.IDIn(groupIDs...)))
	}
	s, err := util.UserStateToEnt(state)
	if err != nil {
		return nil, 0, err
	}
	query = query.Where(user.StateEQ(s))
	if username != "" {
		query = query.Where(user.UsernameContains(username))
	}
	res, err := query.Offset(int((current - 1) * size)).Limit(int(size)).WithGroups().All(ctx)
	if err != nil {
		log.Debug("err: ", err)
		return nil, 0, err
	}
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	rv := make([]*biz.User, 0)
	for _, u := range res {
		g := make([]int64, 0, len(u.Edges.Groups))
		for _, v := range u.Edges.Groups {
			g = append(g, v.ID)
		}
		rv = append(rv, &biz.User{
			ID:       u.ID,
			Username: u.Username,
			GroupIds: g,
			State:    util.UserStateToInt(u.State),
		})
	}
	return rv, int64(total), nil
}
func (r *userRepo) BatchSave(ctx context.Context, users []*biz.UserCreate) ([]string, error) {
	createdUsers, err := r.data.db.User.MapCreateBulk(users, func(c *ent.UserCreate, i int) {
		c.SetUsername(users[i].Username).
			SetSalt(users[i].Salt).
			SetPassword(users[i].Password).
			AddGroupIDs(users[i].GroupIds...)
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
