package biz

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type User struct {
	ID        int64
	Username  string
	Password  string
	Salt      string
	Status    int16
	GroupID   int64
	GroupName string
}

// GreeterRepo is a Greater repo.
type AuthRepo interface {
	FindByUserName(context.Context, string) (*User, error)
}

// GreeterUsecase is a Greeter usecase.
type AuthUsecase struct {
	repo AuthRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AuthUsecase) Login(ctx context.Context, username string, password string) (*User, error) {
	user, err := uc.repo.FindByUserName(ctx, username)
	if err != nil {
		return nil, err
	}
	if verifyPassword(password, user.Salt, user.Password) {
		return user, nil
	}
	return nil, errors.New("password is wrong")

}

func generateMD5Password(password string, salt string) string {
	hash := md5.New()
	hash.Write([]byte(password + salt))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func verifyPassword(password string, salt string, hash string) bool {
	return generateMD5Password(password, salt) == hash
}
