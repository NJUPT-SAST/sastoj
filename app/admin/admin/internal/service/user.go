package service

import (
	"context"
	_ "crypto/md5"
	_ "encoding/hex"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/app/admin/admin/internal/biz"
)

func (s *AdminService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	rv, err := s.uc.CreateUser(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
		GroupIds: req.GroupIds,
		State:    0,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id: rv.ID,
	}, nil
}
func (s *AdminService) BatchCreateUser(ctx context.Context, req *pb.BatchCreateUserRequest) (*pb.BatchCreateUserReply, error) {
	resMap, err := s.uc.BatchSave(ctx, int32(req.Number), req.GroupIds)
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
func (s *AdminService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	rv, err := s.uc.UpdateUser(ctx, &biz.User{
		ID:       req.Id,
		Username: req.Username,
		State:    int16(req.State),
		GroupIds: req.GroupIds,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserReply{
		Success: rv,
	}, nil
}
func (s *AdminService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *AdminService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *AdminService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	rv, err := s.uc.ListUser(ctx, req.Current, req.Size, req.GroupIds, req.Username, int16(req.State))
	if err != nil {
		return nil, err
	}
	var users []*pb.ListUserReply_User
	for _, v := range rv {
		users = append(users, &pb.ListUserReply_User{
			Id:       v.ID,
			Username: v.Username,
			GroupIds: v.GroupIds,
			State:    int32(v.State),
		})
	}
	return &pb.ListUserReply{
		Users: users,
	}, nil
}
