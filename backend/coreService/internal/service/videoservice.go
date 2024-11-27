package service

import (
	"context"
	"time"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
	"github.com/LXJ0000/tok/backend/coreService/internal/biz"
)

type VideoServiceService struct {
	pb.UnimplementedVideoServiceServer
	videoUc biz.VideoUsecase
}

func NewVideoServiceService() *VideoServiceService {
	return &VideoServiceService{}
}

func (s *VideoServiceService) FeedShortVideo(ctx context.Context, req *pb.FeedShortVideoRequest) (*pb.FeedShortVideoResponse, error) {
	vs, err := s.videoUc.Feed(ctx, req.UserId, req.LatestTime, int(req.FeedNum))
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.Video, 0, len(vs))
	for _, v := range vs {
		resp = append(resp, biz2Pb(v))
	}
	return &pb.FeedShortVideoResponse{
		Videos: resp,
		Meta:   SuccessMeta,
	}, nil
}

func (s *VideoServiceService) GetVideoById(ctx context.Context, req *pb.GetVideoByIdRequest) (*pb.GetVideoByIdResponse, error) {
	v, err := s.videoUc.GetVideo(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return &pb.GetVideoByIdResponse{
		Video: biz2Pb(v),
		Meta:  SuccessMeta,
	}, nil
}

func (s *VideoServiceService) PublishVideo(ctx context.Context, req *pb.PublishVideoRequest) (*pb.PublishVideoResponse, error) {
	v := &biz.Video{
		Title:       req.Title,
		CoverURL:    req.CoverUrl,
		VideoURL:    req.PlayUrl,
		Description: req.Description,
		UserID:      req.UserId,
	}
	if err := s.videoUc.CreateVideo(ctx, v); err != nil {
		return nil, err
	}
	return &pb.PublishVideoResponse{
		VideoId: v.ID,
		Meta:    SuccessMeta,
	}, nil
}

func (s *VideoServiceService) ListPublishedVideo(ctx context.Context, req *pb.ListPublishedVideoRequest) (*pb.ListPublishedVideoResponse, error) {
	vs, err := s.videoUc.GetVideoByUserID(ctx, req.UserId, req.LatestTime , int(req.Pagination.Size))
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.Video, 0, len(vs))
	for _, v := range vs {
		resp = append(resp, biz2Pb(v))
	}
	return &pb.ListPublishedVideoResponse{
		Videos: resp,
		Meta:   SuccessMeta,
	}, nil
}

func (s *VideoServiceService) GetVideoByIdList(ctx context.Context, req *pb.GetVideoByIdListRequest) (*pb.GetVideoByIdListResponse, error) {
	vs, err := s.videoUc.GetVideoByIdList(ctx, req.VideoIdList)
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.Video, 0, len(vs))
	for _, v := range vs {
		resp = append(resp, biz2Pb(v))
	}
	return &pb.GetVideoByIdListResponse{
		Videos: resp,
		Meta:   SuccessMeta,
	}, nil
}

func biz2Pb(v *biz.Video) *pb.Video {
	return &pb.Video{
		Id:    v.ID,
		Title: v.Title,
		Author: &pb.Author{
			Id:   v.Author.ID,
			Name: v.Author.Name,
		},
		PlayUrl:       v.VideoURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: v.LikeCount,
		CollectCount:  v.CollectCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		IsCollect:     v.IsCollect,
		Description:   v.Description,
		UploadTime:    v.UpdateTime.Format(time.DateTime),
	}
}
