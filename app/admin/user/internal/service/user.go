package service

import (
	"context"
	_ "crypto/md5"
	_ "encoding/hex"
	pb "sastoj/api/sastoj/admin/user/service/v1"
	"sastoj/app/admin/user/internal/biz"
	"sastoj/pkg/util"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(user *biz.UserUsecase) *UserService {
	return &UserService{
		uc: user,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	var salt = util.GenerateRandomString(5, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var md5Password = util.GenerateMD5Password(req.Password, salt)
	//var b = verifyPassword(req.Password, salt, md5Password)
	//print(b)
	rv, err := s.uc.CreateUser(ctx, &biz.User{
		Username: req.Username,
		Password: md5Password,
		GroupID:  req.GroupId,
		Salt:     salt,
		Status:   0,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id: rv.ID,
	}, nil
}
func (s *UserService) BatchCreateUser(ctx context.Context, req *pb.BatchCreateUserRequest) (*pb.BatchCreateUserReply, error) {

	resMap, err := s.uc.BatchSave(ctx, int32(req.Number), req.GroupId)
	if err != nil {
		return nil, err
	}

	var users []*pb.BatchCreateUserReply_User
	for k, v := range resMap {
		users = append(users, &pb.BatchCreateUserReply_User{
			Username: k,
			Password: v,
		})
	}
	return &pb.BatchCreateUserReply{
		Users: users,
	}, nil

}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	rv, err := s.uc.UpdateUser(ctx, &biz.User{
		ID:       req.Id,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserReply{
		Success: rv,
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	rv, err := s.uc.ListUser(ctx, req.Current, req.Size)
	if err != nil {
		return nil, err
	}
	var users []*pb.ListUserReply_User
	for _, v := range rv {
		users = append(users, &pb.ListUserReply_User{
			Id:       v.ID,
			Username: v.Username,
		})
	}
	return &pb.ListUserReply{
		Users: users,
	}, nil
}
