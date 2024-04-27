package service

import (
	"context"
	"crypto/md5"
	_ "crypto/md5"
	"encoding/hex"
	_ "encoding/hex"
	"github.com/go-kratos/kratos/v2/log"
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
	var resMap = make(map[string]string)
	for i := 0; i < int(req.Number); i++ {
		var username = generateRandomString(8, "")
		var salt = generateRandomString(5, "")
		var password = generateRandomString(8, "")
		var md5Password = generateMD5Password(password, salt)
		_, err := s.uc.CreateUser(ctx, &biz.User{
			Username: username,
			Password: md5Password,
			GroupID:  req.GroupId,
			Salt:     salt,
			Status:   0,
		})
		if err != nil {
			log.Debug("BatchCreateUser error: ", err)
			continue
		}
		resMap[username] = password
	}
	var users []*pb.BatchCreateUserReply_User
	for k, v := range resMap {
		users = append(users, &pb.BatchCreateUserReply_User{
			Username: v,
			Password: k,
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

func generateRandomString(length int, charset string) string {
	if charset == "" {
		charset = "abcdefghjkmnopqrstwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
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
