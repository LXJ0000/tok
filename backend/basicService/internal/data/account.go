package data

import (
	"context"

	"github.com/LXJ0000/tok/backend/basicService/internal/biz"
	"github.com/LXJ0000/tok/backend/basicService/internal/data/model"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepo) Save(ctx context.Context, g *biz.Account) error {
	var account model.Account
	copier.Copy(&account, g)
	if err := r.data.db.Model(&model.Account{}).Create(&account).Error; err != nil {
		return err
	}
	copier.Copy(g, &account)
	return nil
}

func (r *accountRepo) Update(ctx context.Context, g *biz.Account) error {
	var account model.Account
	copier.Copy(&account, g)
	return r.data.db.Model(&model.Account{}).Where("id = ?", g.ID).Updates(&account).Error
}

func (r *accountRepo) FindByID(ctx context.Context, id int64) (*biz.Account, error) {
	var account model.Account
	if err := r.data.db.Model(&model.Account{}).Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	var g biz.Account
	copier.Copy(&g, &account)
	return &g, nil
}

func (r *accountRepo) FindByMobile(ctx context.Context, mobile string) (*biz.Account, error) {
	var account model.Account
	if err := r.data.db.Model(&model.Account{}).Where("mobile = ?", mobile).First(&account).Error; err != nil {
		return nil, err
	}
	var g biz.Account
	copier.Copy(&g, &account)
	return &g, nil
}

func (r *accountRepo) FindByEmail(ctx context.Context, email string) (*biz.Account, error) {
	var account model.Account
	if err := r.data.db.Model(&model.Account{}).Where("email = ?", email).First(&account).Error; err != nil {
		return nil, err
	}
	var g biz.Account
	copier.Copy(&g, &account)
	return &g, nil
}
