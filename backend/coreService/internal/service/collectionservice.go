package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
)

type CollectionServiceService struct {
	pb.UnimplementedCollectionServiceServer
}

func NewCollectionServiceService() *CollectionServiceService {
	return &CollectionServiceService{}
}

func (s *CollectionServiceService) CreateCollection(ctx context.Context, req *pb.CreateCollectionRequest) (*pb.CreateCollectionResponse, error) {
	return &pb.CreateCollectionResponse{}, nil
}
func (s *CollectionServiceService) GetCollectionById(ctx context.Context, req *pb.GetCollectionByIdRequest) (*pb.GetCollectionByIdResponse, error) {
	return &pb.GetCollectionByIdResponse{}, nil
}
func (s *CollectionServiceService) RemoveCollection(ctx context.Context, req *pb.RemoveCollectionRequest) (*pb.RemoveCollectionResponse, error) {
	return &pb.RemoveCollectionResponse{}, nil
}
func (s *CollectionServiceService) ListCollection(ctx context.Context, req *pb.ListCollectionRequest) (*pb.ListCollectionResponse, error) {
	return &pb.ListCollectionResponse{}, nil
}
func (s *CollectionServiceService) UpdateCollection(ctx context.Context, req *pb.UpdateCollectionRequest) (*pb.UpdateCollectionResponse, error) {
	return &pb.UpdateCollectionResponse{}, nil
}
func (s *CollectionServiceService) AddVideo2Collection(ctx context.Context, req *pb.AddVideo2CollectionRequest) (*pb.AddVideo2CollectionResponse, error) {
	return &pb.AddVideo2CollectionResponse{}, nil
}
func (s *CollectionServiceService) RemoveVideoFromCollection(ctx context.Context, req *pb.RemoveVideoFromCollectionRequest) (*pb.RemoveVideoFromCollectionResponse, error) {
	return &pb.RemoveVideoFromCollectionResponse{}, nil
}
func (s *CollectionServiceService) ListCollectionVideo(ctx context.Context, req *pb.ListCollectionVideoRequest) (*pb.ListCollectionVideoResponse, error) {
	return &pb.ListCollectionVideoResponse{}, nil
}
func (s *CollectionServiceService) IsCollected(ctx context.Context, req *pb.IsCollectedRequest) (*pb.IsCollectedResponse, error) {
	return &pb.IsCollectedResponse{}, nil
}
func (s *CollectionServiceService) CountCollect4Video(ctx context.Context, req *pb.CountCollect4VideoRequest) (*pb.CountCollect4VideoResponse, error) {
	return &pb.CountCollect4VideoResponse{}, nil
}
