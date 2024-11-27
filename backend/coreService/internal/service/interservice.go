package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
)

type InterServiceService struct {
	pb.UnimplementedInterServiceServer
}

func NewInterServiceService() *InterServiceService {
	return &InterServiceService{}
}

func (s *InterServiceService) IncrReadCount(ctx context.Context, req *pb.IncrReadCountRequest) (*pb.IncrReadCountResponse, error) {
	return &pb.IncrReadCountResponse{}, nil
}

func (s *InterServiceService) Like(ctx context.Context, req *pb.LikeRequest) (*pb.LikeResponse, error) {
	return &pb.LikeResponse{}, nil
}

func (s *InterServiceService) CancelLike(ctx context.Context, req *pb.CancelLikeRequest) (*pb.CancelLikeResponse, error) {
	return &pb.CancelLikeResponse{}, nil
}

func (s *InterServiceService) Collect(ctx context.Context, req *pb.CollectRequest) (*pb.CollectResponse, error) {
	return &pb.CollectResponse{}, nil
}

func (s *InterServiceService) CancelCollect(ctx context.Context, req *pb.CancelCollectRequest) (*pb.CancelCollectResponse, error) {
	return &pb.CancelCollectResponse{}, nil
}

func (s *InterServiceService) Stat(ctx context.Context, req *pb.StatRequest) (*pb.StatResponse, error) {
	return &pb.StatResponse{}, nil
}
