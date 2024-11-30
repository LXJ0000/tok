package data

import (
	"context"
	"github.com/LXJ0000/tok/backend/basicService/internal/data/model"
	"github.com/jinzhu/copier"

	"github.com/LXJ0000/tok/backend/basicService/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type fileRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileRepo .
func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *fileRepo) Save(ctx context.Context, g *biz.File) error {
	var file model.File
	if err := copier.Copy(&file, g); err != nil {
		return err
	}
	if err := r.data.db.WithContext(ctx).Model(&model.File{}).Create(&file).Error; err != nil {
		return err
	}
	if err := copier.Copy(g, &file); err != nil {
		return err
	}
	return nil
}

func (r *fileRepo) Update(ctx context.Context, g *biz.File) error {
	var account model.File
	if err := copier.Copy(&account, g); err != nil {
		return err
	}
	return r.data.db.WithContext(ctx).Model(&model.File{}).Where("id = ?", g.ID).Updates(&account).Error
}

func (r *fileRepo) FindByID(ctx context.Context, id int64) (*biz.File, error) {
	var account model.File
	if err := r.data.db.WithContext(ctx).Model(&model.File{}).Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	var g biz.File
	if err := copier.Copy(&g, &account); err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *fileRepo) FindUploadedByID(ctx context.Context, id int64) (*biz.File, error) {
	var account model.File
	if err := r.data.db.WithContext(ctx).Model(&model.File{}).
		Where("id = ? and uploaded = true", id).
		First(&account).Error; err != nil {
		return nil, err
	}
	var g biz.File
	if err := copier.Copy(&g, &account); err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *fileRepo) FindUploadedByHash(ctx context.Context, hash string) (*biz.File, error) {
	var account model.File
	if err := r.data.db.WithContext(ctx).Model(&model.File{}).
		Where("hash = ? and uploaded = true", hash).
		First(&account).Error; err != nil {
		return nil, err
	}
	var g biz.File
	if err := copier.Copy(&g, &account); err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *fileRepo) DeleteByID(ctx context.Context, id int64) error {
	update := map[string]interface{}{
		"is_deleted": 1,
	}
	return r.data.db.WithContext(ctx).Model(&model.File{}).
		Where("id = ?", id).Updates(update).Error
}
