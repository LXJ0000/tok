package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/gateway/api/gateway/v1"
)

type CommentServiceService struct {
	pb.UnimplementedCommentServiceServer
}

func NewCommentServiceService() *CommentServiceService {
	return &CommentServiceService{}
}

func (s *CommentServiceService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	return &pb.CreateCommentResponse{}, nil
}
func (s *CommentServiceService) RemoveComment(ctx context.Context, req *pb.RemoveCommentRequest) (*pb.RemoveCommentResponse, error) {
	return &pb.RemoveCommentResponse{}, nil
}
func (s *CommentServiceService) ListComment4Video(ctx context.Context, req *pb.ListComment4VideoRequest) (*pb.ListComment4VideoResponse, error) {
	return &pb.ListComment4VideoResponse{}, nil
}
func (s *CommentServiceService) ListChildComment(ctx context.Context, req *pb.ListChildCommentRequest) (*pb.ListChildCommentResponse, error) {
	return &pb.ListChildCommentResponse{}, nil
}
