package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
	"github.com/LXJ0000/tok/backend/coreService/internal/biz"
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
	user := &biz.User{
		AccountID: req.AccountId,
		Mobile:    req.Mobile,
		Email:     req.Email,
	}
	if err := s.userUc.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		UserId: user.ID,
		Meta:   SuccessMeta,
	}, nil
}

func (s *UserServiceService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	u, err := s.userUc.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoResponse{
		User: &pb.User{
			Id:              u.ID,
			Name:            u.Name,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			Mobile:          u.Mobile,
			Email:           u.Email,
		},
		Meta: SuccessMeta,
	}, nil
}

func (s *UserServiceService) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	user := &biz.User{
		ID:              req.UserId,
		Name:            req.Name,
		Avatar:          req.Avatar,
		BackgroundImage: req.BackgroundImage,
		Signature:       req.Signature,
	}
	if _, err := s.userUc.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return &pb.UpdateUserInfoResponse{
		Meta: SuccessMeta,
	}, nil
}

func (s *UserServiceService) GetUserByIdList(ctx context.Context, req *pb.GetUserByIdListRequest) (*pb.GetUserByIdListResponse, error) {
	users, err := s.userUc.GetUserByIdList(ctx, req.UserIdList)
	if err != nil {
		return nil, err
	}
	userList := make([]*pb.User, 0, len(users))
	for _, u := range users {
		userList = append(userList, &pb.User{
			Id:              u.ID,
			Name:            u.Name,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			Mobile:          u.Mobile,
			Email:           u.Email,
		})
	}
	return &pb.GetUserByIdListResponse{
		UserList: userList,
		Meta:     SuccessMeta,
	}, nil
}
