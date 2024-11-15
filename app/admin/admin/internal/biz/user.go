package biz

import (
	"context"
	"sastoj/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is admin not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "admin not found")
)

// User is a User model.
type User struct {
	ID       int64
	Username string
	Password string
	Salt     string
	State    int16
	GroupIds []int64
}
type UserCreate struct {
	Username string
	Password string
	Salt     string
	GroupIds []int64
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*int64, error)
	AddBlackList(context.Context, int64) error
	RemoveBlackList(context.Context, int64) error
	FindByID(context.Context, int64) (*User, error)
	ListPages(ctx context.Context, current int64, size int64, groupIDs []int64, username string, state int16) ([]*User, int64, error)
	BatchSave(ctx context.Context, users []*UserCreate) ([]string, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, u *User) (*User, error) {
	var salt = util.GenerateRandomString(5, "")
	u.Salt = salt
	u.Password = util.GenerateMD5Password(u.Password, salt)
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.Username)
	res, err := uc.repo.Save(ctx, u)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (uc *UserUsecase) UpdateUser(ctx context.Context, u *User) (bool, error) {
	uc.log.WithContext(ctx).Infof("UpdateUser: %v", u.Username)
	rv, err := uc.repo.Update(ctx, u)
	if err != nil || *rv == 0 {
		return false, err
	}
	if u.State != 0 {
		err := uc.repo.AddBlackList(ctx, u.ID)
		if err != nil {
			return false, err
		}
	} else {
		err := uc.repo.RemoveBlackList(ctx, u.ID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func (uc *UserUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
	res, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (uc *UserUsecase) ListUser(ctx context.Context, current int64, size int64, groupIDs []int64, username string, state int16) ([]*User, int64, error) {
	uc.log.WithContext(ctx).Infof("ListUser: %v", current)
	res, total, err := uc.repo.ListPages(ctx, current, size, groupIDs, username, state)
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (uc *UserUsecase) BatchSave(ctx context.Context, number int32, groupIds []int64) (map[string]string, error) {
	users := make([]*UserCreate, 0)
	passwordMap := make(map[string]string)
	for i := 0; i < int(number); i++ {
		var username = "user_" + util.GenerateRandomString(8, "")
		var salt = util.GenerateRandomString(5, "")
		var password = util.GenerateRandomString(8, "")
		var md5Password = util.GenerateMD5Password(password, salt)
		users = append(users, &UserCreate{
			Username: username,
			Salt:     salt,
			Password: md5Password,
			GroupIds: groupIds,
		})
		passwordMap[username] = password
	}
	usernames, err := uc.repo.BatchSave(ctx, users)
	if err != nil {
		return nil, err
	}
	returnedMaps := make(map[string]string)
	for _, v := range usernames {
		//只返回创建成功的账户
		if _, ok := passwordMap[v]; ok {
			returnedMaps[v] = passwordMap[v]
		}
	}
	return returnedMaps, nil
}
