package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/gateway/api/gateway/v1"
)

type ShortVideoCoreVideoServiceService struct {
	pb.UnimplementedShortVideoCoreVideoServiceServer
}

func NewShortVideoCoreVideoServiceService() *ShortVideoCoreVideoServiceService {
	return &ShortVideoCoreVideoServiceService{}
}

func (s *ShortVideoCoreVideoServiceService) PreSign4UploadVideo(ctx context.Context, req *pb.PreSign4UploadVideoRequest) (*pb.PreSign4UploadVideoResponse, error) {
	return &pb.PreSign4UploadVideoResponse{}, nil
}
func (s *ShortVideoCoreVideoServiceService) PreSign4UploadCover(ctx context.Context, req *pb.PreSign4UploadRequest) (*pb.PreSign4UploadResponse, error) {
	return &pb.PreSign4UploadResponse{}, nil
}
func (s *ShortVideoCoreVideoServiceService) ReportFinishUpload(ctx context.Context, req *pb.ReportFinishUploadRequest) (*pb.ReportFinishUploadResponse, error) {
	return &pb.ReportFinishUploadResponse{}, nil
}
func (s *ShortVideoCoreVideoServiceService) ReportVideoFinishUpload(ctx context.Context, req *pb.ReportVideoFinishUploadRequest) (*pb.ReportVideoFinishUploadResponse, error) {
	return &pb.ReportVideoFinishUploadResponse{}, nil
}
func (s *ShortVideoCoreVideoServiceService) FeedShortVideo(ctx context.Context, req *pb.FeedShortVideoRequest) (*pb.FeedShortVideoResponse, error) {
	return &pb.FeedShortVideoResponse{}, nil
}
func (s *ShortVideoCoreVideoServiceService) GetVideoById(ctx context.Context, req *pb.GetVideoByIdRequest) (*pb.GetVideoByIdResponse, error) {
	return &pb.GetVideoByIdResponse{}, nil
}
func (s *ShortVideoCoreVideoServiceService) ListPublishedVideo(ctx context.Context, req *pb.ListPublishedVideoRequest) (*pb.ListPublishedVideoResponse, error) {
	return &pb.ListPublishedVideoResponse{}, nil
}
