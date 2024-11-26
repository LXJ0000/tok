package data

import (
	"context"
	"math"

	"github.com/LXJ0000/tok/backend/coreService/internal/biz"
	"github.com/LXJ0000/tok/backend/coreService/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type videoRepo struct {
	data *Data
	log  *log.Helper
}

// NewVideoRepo .
func NewVideoRepo(data *Data, logger log.Logger) biz.VideoRepo {
	return &videoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *videoRepo) Save(ctx context.Context, g *biz.Video) error {
	var video model.Video
	copier.Copy(&video, g)
	if err := r.data.db.Model(&model.Video{}).Create(&video).Error; err != nil {
		return err
	}
	copier.Copy(g, &video)
	return nil
}

func (r *videoRepo) Update(ctx context.Context, g *biz.Video) error {
	var video model.Video
	copier.Copy(&video, g)
	return r.data.db.Model(&model.Video{}).Where("id = ?", g.ID).Updates(&video).Error
}

func (r *videoRepo) FindByID(ctx context.Context, id int64) (*biz.Video, error) {
	var video model.Video
	if err := r.data.db.Model(&model.Video{}).Where("id = ?", id).First(&video).Error; err != nil {
		return nil, err
	}
	var g biz.Video
	copier.Copy(&g, &video)
	return &g, nil
}

func (r *videoRepo) FindByIDList(ctx context.Context, ids []int64) ([]*biz.Video, error) {
	var videos []*model.Video
	if err := r.data.db.Model(&model.Video{}).Where("id in (?)", ids).Find(&videos).Error; err != nil {
		return nil, err
	}
	var greeters []*biz.Video
	for _, video := range videos {
		var g biz.Video
		copier.Copy(&g, video)
		greeters = append(greeters, &g)
	}
	return greeters, nil
}

func (r *videoRepo) FindByUserID(ctx context.Context, userID int64, lastID int64, limit int) ([]*biz.Video, error) {
	var videos []*model.Video
	if lastID <= 0 {
		lastID = math.MaxInt64
	}
	if err := r.data.db.Model(&model.Video{}).Where("user_id = ? and id < ?", userID, lastID).Limit(limit).Order("id desc").Find(&videos).Error; err != nil {
		return nil, err
	}
	var greeters []*biz.Video
	for _, video := range videos {
		var g biz.Video
		copier.Copy(&g, video)
		greeters = append(greeters, &g)
	}
	return greeters, nil
}

func (r *videoRepo) Find(ctx context.Context, lastTime int64, limit int) ([]*biz.Video, error) {
	var videos []*model.Video
	if err := r.data.db.Model(&model.Video{}).Where("create_time < ?", lastTime).Limit(limit).Order("id desc").Find(&videos).Error; err != nil {
		return nil, err
	}
	var greeters []*biz.Video
	for _, video := range videos {
		var g biz.Video
		copier.Copy(&g, video)
		greeters = append(greeters, &g)
	}
	return greeters, nil
}
