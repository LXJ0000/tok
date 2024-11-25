package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Account is a Account model.
type Account struct {
	ID         int64     ` json:"id"`          // 账户ID
	Mobile     string    ` json:"mobile"`      // 手机号
	Email      string    ` json:"email"`       // 电子邮件
	Password   string    ` json:"password"`    // 密码
	Salt       string    ` json:"salt"`        // 密码盐
	Number     string    ` json:"number"`      // 号码
	IsDeleted  bool      ` json:"is_deleted"`  // 是否删除
	CreateTime time.Time ` json:"create_time"` // 创建时间
	UpdateTime time.Time ` json:"update_time"` // 更新时间
}

// AccountRepo is a Greater repo.
type AccountRepo interface {
	Save(context.Context, *Account) error
	Update(context.Context, *Account) (*Account, error)
	FindByID(context.Context, int64) (*Account, error)
}

// AccountUsecase is a Account usecase.
type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

// NewAccountUsecase new a Account usecase.
func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateAccount creates a Account, and returns the new Account.
func (uc *AccountUsecase) CreateAccount(ctx context.Context, g *Account) error {
	return uc.repo.Save(ctx, g)
}
