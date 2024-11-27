package data

import (
	"context"
	"errors"

	domain "github.com/LXJ0000/tok/backend/coreService/internal/biz"
	"github.com/LXJ0000/tok/backend/coreService/internal/data/model"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/go-kratos/kratos/v2/log"
)

type interRepo struct {
	data *Data
	log  *log.Helper
}

// NewInterRepo .
func NewInterRepo(data *Data, logger log.Logger) domain.InterRepo {
	return &interRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *interRepo) IncrReadCount(c context.Context, biz string, id int64) error {
	update := map[string]interface{}{
		"read_cnt": gorm.Expr("`read_cnt` + 1"),
	}
	create := &model.Interaction{
		BizID:   id,
		Biz:     biz,
		ReadCnt: 1,
	}
	err := r.data.db.WithContext(c).Model(&model.Interaction{}).Clauses(clause.OnConflict{
		//Columns:  // mysql 可以不写
		DoUpdates: clause.Assignments(update),
	}).Create(create).Error
	if err != nil {
		return err
	}
	// 数据库操作成功即认为业务处理成功
	return nil
}

func (r *interRepo) Like(c context.Context, biz string, bizID, userID int64) error {
	// upsert interaction 表
	updateInteraction := map[string]interface{}{
		"like_cnt": gorm.Expr("`like_cnt` + 1"),
	}
	createInteraction := &model.Interaction{
		BizID:   bizID,
		Biz:     biz,
		LikeCnt: 1,
	}
	// upsert user_like 表
	updateUserLike := map[string]interface{}{
		"status": true,
	}
	createUserLike := &model.UserLike{
		BizID:  bizID,
		Biz:    biz,
		UserID: userID,
		Status: true,
	}
	fn := func(tx *gorm.DB) error {
		if err := r.data.db.WithContext(c).Model(&model.Interaction{}).Clauses(clause.OnConflict{
			//Columns:  // mysql 可以不写
			DoUpdates: clause.Assignments(updateInteraction),
		}).Create(createInteraction).Error; err != nil {
			return err
		}
		return r.data.db.WithContext(c).Model(&model.UserLike{}).Clauses(clause.OnConflict{
			//Columns:  // mysql 可以不写
			DoUpdates: clause.Assignments(updateUserLike),
		}).Create(createUserLike).Error
	}
	if err := r.data.db.WithContext(c).Transaction(fn); err != nil {
		return err
	}
	return nil
}

func (r *interRepo) CancelLike(c context.Context, biz string, bizID, userID int64) error {
	// upsert interaction 表
	updateInteraction := map[string]interface{}{
		"like_cnt": gorm.Expr("`like_cnt` - 1"),
	}
	createInteraction := &model.Interaction{
		BizID: bizID,
		Biz:   biz,
	}
	// upsert user_like 表
	updateUserLike := map[string]interface{}{
		"status": false,
	}
	createUserLike := &model.UserLike{
		BizID:  bizID,
		Biz:    biz,
		UserID: userID,
	}
	fn := func(tx *gorm.DB) error {
		if err := r.data.db.WithContext(c).Model(&model.Interaction{}).Clauses(clause.OnConflict{
			//Columns:  // mysql 可以不写
			DoUpdates: clause.Assignments(updateInteraction),
		}).Create(createInteraction).Error; err != nil {
			return err
		}
		return r.data.db.WithContext(c).Model(&model.UserLike{}).Clauses(clause.OnConflict{
			//Columns:  // mysql 可以不写
			DoUpdates: clause.Assignments(updateUserLike),
		}).Create(createUserLike).Error
	}
	if err := r.data.db.WithContext(c).Transaction(fn); err != nil {
		return err
	}
	return nil
}

func (r *interRepo) Collect(c context.Context, biz string, bizID, userID, collectionID int64) error {
	// upsert interaction 表
	updateInteraction := map[string]interface{}{
		"collect_cnt": gorm.Expr("`collect_cnt` + 1"),
	}
	createInteraction := &model.Interaction{
		BizID:      bizID,
		Biz:        biz,
		CollectCnt: 1,
	}
	// upsert collect_like 表
	updateUserCollect := map[string]interface{}{
		"status": true,
	}
	createUserCollect := &model.UserCollect{
		BizID:  bizID,
		Biz:    biz,
		UserID: userID,
		Status: true,
	}
	fn := func(tx *gorm.DB) error {
		if err := r.data.db.WithContext(c).Model(&model.Interaction{}).Clauses(clause.OnConflict{
			//Columns:  // mysql 可以不写
			DoUpdates: clause.Assignments(updateInteraction),
		}).Create(createInteraction).Error; err != nil {
			return err
		}
		return r.data.db.WithContext(c).Model(&model.UserCollect{}).Clauses(clause.OnConflict{
			//Columns:  // mysql 可以不写
			DoUpdates: clause.Assignments(updateUserCollect),
		}).Create(createUserCollect).Error
	}
	if err := r.data.db.WithContext(c).Transaction(fn); err != nil {
		return err
	}
	return nil
}

func (r *interRepo) CancelCollect(c context.Context, biz string, bizID, userID, collectionID int64) error {
	// upsert interaction 表
	updateInteraction := map[string]interface{}{
		"collect_cnt": gorm.Expr("`collect_cnt` - 1"),
	}
	createInteraction := &model.Interaction{
		BizID: bizID,
		Biz:   biz,
	}
	// upsert user_like 表
	updateUserCollect := map[string]interface{}{
		"status": false,
	}
	createUserCollect := &model.UserCollect{
		BizID:  bizID,
		Biz:    biz,
		UserID: userID,
	}
	fn := func(tx *gorm.DB) error {
		if err := r.data.db.WithContext(c).Model(&model.Interaction{}).Clauses(clause.OnConflict{
			//Columns:  // mysql 可以不写
			DoUpdates: clause.Assignments(updateInteraction),
		}).Create(createInteraction).Error; err != nil {
			return err
		}
		return r.data.db.WithContext(c).Model(&model.UserCollect{}).Clauses(clause.OnConflict{
			//Columns:  // mysql 可以不写
			DoUpdates: clause.Assignments(updateUserCollect),
		}).Create(createUserCollect).Error
	}
	if err := r.data.db.WithContext(c).Transaction(fn); err != nil {
		return err
	}
	return nil
}

func (r *interRepo) Stat(c context.Context, biz string, bizID, userID int64) (*domain.InterStat, error) {
	var isLike, isCollect bool
	var err error
	var interaction model.Interaction
	eg := errgroup.Group{}
	eg.Go(func() error {
		isLike, err = r.isLike(c, biz, bizID, userID)
		if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	})
	eg.Go(func() error {
		isCollect, err = r.isCollect(c, biz, bizID, userID)
		if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	})
	eg.Go(func() error {
		var interaction model.Interaction
		filter := map[string]interface{}{
			"biz_id": bizID,
			"biz":    biz,
		}

		if err := r.data.db.WithContext(c).Model(&model.Interaction{}).Where(filter).First(&interaction).Error; err != nil {
			return nil
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return &domain.InterStat{
		LikeCnt:    interaction.LikeCnt,
		CollectCnt: interaction.CollectCnt,
		ReadCnt:    interaction.ReadCnt,
		IsLike:     isLike,
		IsCollect:  isCollect,
	}, nil
}

func (r *interRepo) isLike(c context.Context, biz string, bizID, userID int64) (bool, error) {
	var item model.UserLike
	filter := map[string]interface{}{
		"user_id": userID,
		"biz_id":  bizID,
		"biz":     biz,
	}
	if err := r.data.db.WithContext(c).Model(&model.UserLike{}).Where(filter).First(&item).Error; err != nil {
		return false, err
	}
	return item.Status, nil
}

func (r *interRepo) isCollect(c context.Context, biz string, bizID, userID int64) (bool, error) {
	var item model.UserCollect
	filter := map[string]interface{}{
		"user_id": userID,
		"biz_id":  bizID,
		"biz":     biz,
	}
	if err := r.data.db.WithContext(c).Model(&model.UserCollect{}).Where(filter).First(&item).Error; err != nil {
		return false, err
	}
	return item.Status, nil
}
