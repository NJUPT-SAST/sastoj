package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"sastoj/app/public/auth/internal/biz"

	pb "sastoj/api/sastoj/public/auth/service/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	au *biz.AuthUsecase
}

func NewAuthService(au *biz.AuthUsecase) *AuthService {
	return &AuthService{au: au}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := s.au.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	jwtToken, err := generateJWT(user)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: jwtToken,
	}, nil

}
func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}

func generateJWT(user *biz.User) (string, error) {
	// 创建一个新的 Token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置声明（Claims）
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.ID           // 设置主题
	claims["groupId"] = user.GroupID     // 设置发布者
	claims["groupName"] = user.GroupName // 设置发布者
	// 设置密钥
	secret := []byte("sastoj") // 替换为实际的密钥

	// 使用密钥对 Token 进行签名，生成最终的 JWT
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
