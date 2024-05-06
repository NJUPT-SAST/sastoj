package biz

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"sastoj/app/public/auth/internal/conf"
	"time"
)

// User is a Greeter model.
type User struct {
	ID        int64
	Username  string
	Password  string
	Salt      string
	Status    int16
	GroupID   int64
	GroupName string
}
type MyCustomClaims struct {
	UserId     int64   `json:"user_id"`
	GroupId    int64   `json:"group_id"`
	GroupName  string  `json:"group_name"`
	ContestIds []int64 `json:"contest_ids"`
	jwt.RegisteredClaims
}

// AuthRepo is a Greater repo.
type AuthRepo interface {
	FindByUserName(context.Context, string) (*User, error)
}

type jwtConfig struct {
	Secret     string
	Expiration time.Duration
	Issuer     string
}

// AuthUsecase is a Greeter usecase.
type AuthUsecase struct {
	repo    AuthRepo
	log     *log.Helper
	jwtConf jwtConfig
}

// NewAuthUsecase new a Greeter usecase.
func NewAuthUsecase(repo AuthRepo, logger log.Logger, c *conf.Jwt) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger), jwtConf: jwtConfig{
		Secret:     c.Secret,
		Expiration: time.Duration(c.Expiration) * time.Hour,
		Issuer:     c.Issuer},
	}
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
func (uc *AuthUsecase) GenerateJWT(user User, contestIds []int64) (string, error) {
	// 设置声明（Claims）
	registeredClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.jwtConf.Expiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    uc.jwtConf.Issuer,
		Subject:   "a",
		ID:        "1",
	}
	claims := MyCustomClaims{
		UserId:           user.ID,
		GroupId:          user.GroupID,
		GroupName:        user.GroupName,
		ContestIds:       contestIds,
		RegisteredClaims: registeredClaims,
	}
	// 使用密钥对 Token 进行签名，生成最终的 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(uc.jwtConf.Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
