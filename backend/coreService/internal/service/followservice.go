package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
)

type FollowServiceService struct {
	pb.UnimplementedFollowServiceServer
}

func NewFollowServiceService() *FollowServiceService {
	return &FollowServiceService{}
}

func (s *FollowServiceService) AddFollow(ctx context.Context, req *pb.AddFollowRequest) (*pb.AddFollowResponse, error) {
	return &pb.AddFollowResponse{}, nil
}
func (s *FollowServiceService) RemoveFollow(ctx context.Context, req *pb.RemoveFollowRequest) (*pb.RemoveFollowResponse, error) {
	return &pb.RemoveFollowResponse{}, nil
}
func (s *FollowServiceService) ListFollowing(ctx context.Context, req *pb.ListFollowingRequest) (*pb.ListFollowingResponse, error) {
	return &pb.ListFollowingResponse{}, nil
}
func (s *FollowServiceService) IsFollowing(ctx context.Context, req *pb.IsFollowingRequest) (*pb.IsFollowingResponse, error) {
	return &pb.IsFollowingResponse{}, nil
}
func (s *FollowServiceService) CountFollow(ctx context.Context, req *pb.CountFollowRequest) (*pb.CountFollowResponse, error) {
	return &pb.CountFollowResponse{}, nil
}
