package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/basicService/api/basicService/v1"
)

type AuthServiceService struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceService() *AuthServiceService {
	return &AuthServiceService{}
}

func (s *AuthServiceService) CreateVerificationCode(ctx context.Context, req *pb.CreateVerificationCodeRequest) (*pb.CreateVerificationCodeResponse, error) {
	return &pb.CreateVerificationCodeResponse{}, nil
}
func (s *AuthServiceService) ValidateVerificationCode(ctx context.Context, req *pb.ValidateVerificationCodeRequest) (*pb.ValidateVerificationCodeResponse, error) {
	return &pb.ValidateVerificationCodeResponse{}, nil
}
