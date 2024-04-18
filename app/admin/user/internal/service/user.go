package service

import (
	"context"
	"crypto/md5"
	_ "crypto/md5"
	"encoding/hex"
	_ "encoding/hex"
	"math/rand"
	pb "sastoj/api/sastoj/admin/user/service/v1"
	"sastoj/app/admin/user/internal/biz"
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
	var salt = generateRandomString(5, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var md5Password = generateMD5Password(req.Password, salt)
	var b = verifyPassword(req.Password, salt, md5Password)
	print(b)
	rv, err := s.uc.CreateUser(ctx, &biz.User{
		Username: req.Username,
		Password: md5Password,
		GroupID:  int(req.GroupId),
		Salt:     salt,
		State:    1,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id: int64(rv.ID),
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	rv, err := s.uc.UpdateUser(ctx, &biz.User{
		ID:       int(req.Id),
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
	rv, err := s.uc.ListUser(ctx, int(req.Current), int(req.Size))
	if err != nil {
		return nil, err
	}
	var users []*pb.ListUserReply_User
	for _, v := range rv {
		users = append(users, &pb.ListUserReply_User{
			Id:       int64(v.ID),
			Username: v.Username,
		})
	}
	return &pb.ListUserReply{
		Users: users,
	}, nil
}

func generateRandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
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
