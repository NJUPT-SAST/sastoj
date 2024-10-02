package service

import (
	"context"
	pb "sastoj/api/sastoj/public/auth/service/v1"
	"sastoj/app/public/auth/internal/biz"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	au *biz.AuthUsecase
}

func NewAuthService(au *biz.AuthUsecase) *AuthService {
	return &AuthService{au: au}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	jwtToken, err := s.au.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: jwtToken,
	}, nil

}

// func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
// 	return &pb.LogoutReply{}, nil
// }
