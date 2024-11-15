package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/gateway/api/gateway/v1"
)

type FavoriteServiceService struct {
	pb.UnimplementedFavoriteServiceServer
}

func NewFavoriteServiceService() *FavoriteServiceService {
	return &FavoriteServiceService{}
}

func (s *FavoriteServiceService) AddFavorite(ctx context.Context, req *pb.AddFavoriteRequest) (*pb.AddFavoriteResponse, error) {
	return &pb.AddFavoriteResponse{}, nil
}
func (s *FavoriteServiceService) RemoveFavorite(ctx context.Context, req *pb.RemoveFavoriteRequest) (*pb.RemoveFavoriteResponse, error) {
	return &pb.RemoveFavoriteResponse{}, nil
}
func (s *FavoriteServiceService) ListFavoriteVideo(ctx context.Context, req *pb.ListFavoriteVideoRequest) (*pb.ListFavoriteVideoResponse, error) {
	return &pb.ListFavoriteVideoResponse{}, nil
}
