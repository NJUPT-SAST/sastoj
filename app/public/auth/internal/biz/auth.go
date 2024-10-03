package biz

import (
	"context"
	"errors"
	"sastoj/app/public/auth/internal/conf"
	"sastoj/pkg/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

// User is a user model.
type User struct {
	ID        int64
	Username  string
	Password  string
	Salt      string
	Status    int16
	GroupIds  []int64
	GroupName string
}

// MyCustomClaims is a custom claim for jwt.
type MyCustomClaims struct {
	UserId    int64   `json:"user_id"`
	GroupIds  []int64 `json:"group_ids"`
	GroupName string  `json:"group_name"`
	jwt.RegisteredClaims
}

// AuthRepo is a user and contest repo.
type AuthRepo interface {
	FindUserByName(context.Context, string) (*User, error)
}

type jwtConfig struct {
	secret     string
	expiration time.Duration
	issuer     string
}

// AuthUsecase is a Greeter usecase.
type AuthUsecase struct {
	repo    AuthRepo
	log     *log.Helper
	jwtConf jwtConfig
}

// NewAuthUsecase new a Greeter usecase.
func NewAuthUsecase(repo AuthRepo, logger log.Logger, c *conf.Jwt) *AuthUsecase {
	var exp int32 = 2
	if c.Expiration != 0 {
		exp = c.Expiration
	}
	var issuer = "sastoj"
	if c.Issuer != "" {
		issuer = c.Issuer
	}
	return &AuthUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		jwtConf: jwtConfig{
			secret:     c.Secret,
			expiration: time.Duration(exp) * time.Hour,
			issuer:     issuer,
		},
	}
}

func (uc *AuthUsecase) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := uc.repo.FindUserByName(ctx, username)
	if err != nil {
		return "", errors.New("user or password is wrong")
	}
	if !util.VerifyPassword(password, user.Salt, user.Password) {
		return "", errors.New("user or password is wrong")
	}
	j, err := uc.generateJWT(user)
	if err != nil {
		return "", err
	}
	return j, nil
}

func (uc *AuthUsecase) generateJWT(user *User) (string, error) {
	// 设置声明（Claims）
	registeredClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.jwtConf.expiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    uc.jwtConf.issuer,
		Subject:   "a",
		ID:        "1",
	}
	claims := MyCustomClaims{
		UserId:           user.ID,
		GroupIds:         user.GroupIds,
		GroupName:        user.GroupName,
		RegisteredClaims: registeredClaims,
	}
	// 使用密钥对 Token 进行签名，生成最终的 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(uc.jwtConf.secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
