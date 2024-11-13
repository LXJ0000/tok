package service

import (
	"context"

	pb "coreService/api/coreService/v1"
	"coreService/internal/biz"
)

type UserServiceService struct {
	pb.UnimplementedUserServiceServer
	userUc *biz.UserUsecase
}

func NewUserServiceService(userUc *biz.UserUsecase) *UserServiceService {
	return &UserServiceService{
		userUc: userUc,
	}
}

func (s *UserServiceService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}
func (s *UserServiceService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return &pb.GetUserInfoResponse{}, nil
}
func (s *UserServiceService) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	return &pb.UpdateUserInfoResponse{}, nil
}
func (s *UserServiceService) GetUserByIdList(ctx context.Context, req *pb.GetUserByIdListRequest) (*pb.GetUserByIdListResponse, error) {
	return &pb.GetUserByIdListResponse{}, nil
}
