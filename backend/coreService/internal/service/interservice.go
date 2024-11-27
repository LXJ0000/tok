package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
	"github.com/LXJ0000/tok/backend/coreService/internal/biz"
)

type InterServiceService struct {
	pb.UnimplementedInterServiceServer
	interUc *biz.InterUsecase
}

func NewInterServiceService(interUc *biz.InterUsecase) *InterServiceService {
	return &InterServiceService{interUc: interUc}
}

func (s *InterServiceService) IncrReadCount(ctx context.Context, req *pb.IncrReadCountRequest) (*pb.IncrReadCountResponse, error) {
	if err := s.interUc.IncrReadCount(ctx, req.Biz, req.BizId); err != nil {
		return nil, err
	}
	return &pb.IncrReadCountResponse{
		Meta: SuccessMeta,
	}, nil
}

func (s *InterServiceService) Like(ctx context.Context, req *pb.LikeRequest) (*pb.LikeResponse, error) {
	if err := s.interUc.Like(ctx, req.Biz, req.BizId, req.UserId); err != nil {
		return nil, err
	}
	return &pb.LikeResponse{Meta: SuccessMeta}, nil
}

func (s *InterServiceService) CancelLike(ctx context.Context, req *pb.CancelLikeRequest) (*pb.CancelLikeResponse, error) {
	if err := s.interUc.CancelLike(ctx, req.Biz, req.BizId, req.UserId); err != nil {
		return nil, err
	}
	return &pb.CancelLikeResponse{Meta: SuccessMeta}, nil
}

func (s *InterServiceService) Collect(ctx context.Context, req *pb.CollectRequest) (*pb.CollectResponse, error) {
	if err := s.interUc.Collect(ctx, req.Biz, req.BizId, req.UserId, req.CollectionId); err != nil {
		return nil, err
	}
	return &pb.CollectResponse{Meta: SuccessMeta}, nil
}

func (s *InterServiceService) CancelCollect(ctx context.Context, req *pb.CancelCollectRequest) (*pb.CancelCollectResponse, error) {
	if err := s.interUc.CancelCollect(ctx, req.Biz, req.BizId, req.UserId, req.CollectionId); err != nil {
		return nil, err
	}
	return &pb.CancelCollectResponse{Meta: SuccessMeta}, nil
}

func (s *InterServiceService) Stat(ctx context.Context, req *pb.StatRequest) (*pb.StatResponse, error) {
	stat, err := s.interUc.Stat(ctx, req.Biz, req.BizId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.StatResponse{
		LikeCnt:    stat.LikeCnt,
		CollectCnt: stat.CollectCnt,
		ReadCnt:    stat.ReadCnt,
		IsLike:     stat.IsLike,
		IsCollect:  stat.IsCollect,
		Meta:       SuccessMeta,
	}, nil
}
