package data

import (
	"context"

	"github.com/LXJ0000/tok/backend/coreService/internal/biz"
	"github.com/LXJ0000/tok/backend/coreService/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, g *biz.User) error {
	var user model.User
	copier.Copy(&user, g)
	if err := r.data.db.Model(&model.User{}).Create(&user).Error; err != nil {
		return err
	}
	copier.Copy(g, &user)
	return nil
}

func (r *userRepo) Update(ctx context.Context, g *biz.User) error {
	var user model.User
	copier.Copy(&user, g)
	return r.data.db.Model(&model.User{}).Where("id = ?", g.ID).Updates(&user).Error
}

func (r *userRepo) FindByID(ctx context.Context, id int64) (*biz.User, error) {
	var user model.User
	if err := r.data.db.Model(&model.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	var g biz.User
	copier.Copy(&g, &user)
	return &g, nil
}

func (r *userRepo) FindByIDList(ctx context.Context, ids []int64) ([]*biz.User, error) {
	var users []*model.User
	if err := r.data.db.Model(&model.User{}).Where("id in (?)", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	var greeters []*biz.User
	for _, user := range users {
		var g biz.User
		copier.Copy(&g, user)
		greeters = append(greeters, &g)
	}
	return greeters, nil
}
