package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
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
func (s *CommentServiceService) ListChildComment4Comment(ctx context.Context, req *pb.ListChildComment4CommentRequest) (*pb.ListChildComment4CommentResponse, error) {
	return &pb.ListChildComment4CommentResponse{}, nil
}
func (s *CommentServiceService) GetCommentById(ctx context.Context, req *pb.GetCommentByIdRequest) (*pb.GetCommentByIdResponse, error) {
	return &pb.GetCommentByIdResponse{}, nil
}
func (s *CommentServiceService) CountComment4Video(ctx context.Context, req *pb.CountComment4VideoRequest) (*pb.CountComment4VideoResponse, error) {
	return &pb.CountComment4VideoResponse{}, nil
}
func (s *CommentServiceService) CountComment4User(ctx context.Context, req *pb.CountComment4UserRequest) (*pb.CountComment4UserResponse, error) {
	return &pb.CountComment4UserResponse{}, nil
}
