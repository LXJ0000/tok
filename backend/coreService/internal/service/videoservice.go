package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
)

type VideoServiceService struct {
	pb.UnimplementedVideoServiceServer
}

func NewVideoServiceService() *VideoServiceService {
	return &VideoServiceService{}
}

func (s *VideoServiceService) FeedShortVideo(ctx context.Context, req *pb.FeedShortVideoRequest) (*pb.FeedShortVideoResponse, error) {
	return &pb.FeedShortVideoResponse{}, nil
}
func (s *VideoServiceService) GetVideoById(ctx context.Context, req *pb.GetVideoByIdRequest) (*pb.GetVideoByIdResponse, error) {
	return &pb.GetVideoByIdResponse{}, nil
}
func (s *VideoServiceService) PublishVideo(ctx context.Context, req *pb.PublishVideoRequest) (*pb.PublishVideoResponse, error) {
	return &pb.PublishVideoResponse{}, nil
}
func (s *VideoServiceService) ListPublishedVideo(ctx context.Context, req *pb.ListPublishedVideoRequest) (*pb.ListPublishedVideoResponse, error) {
	return &pb.ListPublishedVideoResponse{}, nil
}
func (s *VideoServiceService) GetVideoByIdList(ctx context.Context, req *pb.GetVideoByIdListRequest) (*pb.GetVideoByIdListResponse, error) {
	return &pb.GetVideoByIdListResponse{}, nil
}
