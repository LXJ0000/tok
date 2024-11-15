package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
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
func (s *FavoriteServiceService) ListFavorite(ctx context.Context, req *pb.ListFavoriteRequest) (*pb.ListFavoriteResponse, error) {
	return &pb.ListFavoriteResponse{}, nil
}
func (s *FavoriteServiceService) CountFavorite(ctx context.Context, req *pb.CountFavoriteRequest) (*pb.CountFavoriteResponse, error) {
	return &pb.CountFavoriteResponse{}, nil
}
func (s *FavoriteServiceService) IsFavorite(ctx context.Context, req *pb.IsFavoriteRequest) (*pb.IsFavoriteResponse, error) {
	return &pb.IsFavoriteResponse{}, nil
}
