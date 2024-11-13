package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// User is a User model.
type User struct {
	ID              int64     `json:"id"`
	AccountID       int64     `json:"account_id"`
	Mobile          string    `json:"mobile"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Avatar          string    `json:"avatar"`
	BackgroundImage string    `json:"background_image"`
	Signature       string    `json:"signature"`
	IsDeleted       bool      `json:"is_deleted"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) error
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) error {
	return uc.repo.Save(ctx, g)
}
