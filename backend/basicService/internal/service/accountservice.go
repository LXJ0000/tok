package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/basicService/api/basicService/v1"
)

type AccountServiceService struct {
	pb.UnimplementedAccountServiceServer
}

func NewAccountServiceService() *AccountServiceService {
	return &AccountServiceService{}
}

func (s *AccountServiceService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{}, nil
}
func (s *AccountServiceService) CheckAccount(ctx context.Context, req *pb.CheckAccountRequest) (*pb.CheckAccountResponse, error) {
	return &pb.CheckAccountResponse{}, nil
}
func (s *AccountServiceService) Bind(ctx context.Context, req *pb.BindRequest) (*pb.BindResponse, error) {
	return &pb.BindResponse{}, nil
}
func (s *AccountServiceService) Unbind(ctx context.Context, req *pb.UnbindRequest) (*pb.UnbindResponse, error) {
	return &pb.UnbindResponse{}, nil
}
