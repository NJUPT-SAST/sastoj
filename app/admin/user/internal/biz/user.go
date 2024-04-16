package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User is a User model.
type User struct {
	ID       int
	Username string
	Password string
	Salt     string
	State    int
	GroupID  int
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*int, error)
	FindByID(context.Context, int) (*User, error)
	ListPages(ctx context.Context, current int, size int) ([]*User, error)
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
	return true, nil
}
func (uc *UserUsecase) GetUser(ctx context.Context, id int) (*User, error) {
	res, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (uc *UserUsecase) ListUser(ctx context.Context, current int, size int) ([]*User, error) {
	uc.log.WithContext(ctx).Infof("ListUser: %v", current)
	res, err := uc.repo.ListPages(ctx, current, size)
	if err != nil {
		return nil, err
	}
	return res, nil
}
