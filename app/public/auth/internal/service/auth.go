package service

import (
	"context"
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
	int64s := []int64{1, 2, 3}

	jwtToken, err := s.au.GenerateJWT(*user, int64s)
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
