package biz

import (
	"context"
	"errors"
	"time"

	"github.com/LXJ0000/tok/backend/basicService/internal/infra/utils"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
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
	Update(context.Context, *Account) error
	FindByID(context.Context, int64) (*Account, error)
	FindByMobile(context.Context, string) (*Account, error)
	FindByEmail(context.Context, string) (*Account, error)
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
	if err := uc.checkAccountExist(ctx, *g); err != nil {
		return err
	}
	if !g.isPasswordValid() {
		return errors.New("invalid password with regex")
	}
	if err := g.encryptPassword(); err != nil {
		return err
	}
	g.ID = utils.GenID()

	return uc.repo.Save(ctx, g)
}

func (uc *AccountUsecase) CheckPasswordByMobile(ctx context.Context, mobile, password string) (int64, error) {
	a, err := uc.repo.FindByMobile(ctx, mobile)
	if err != nil {
		return 0, err
	}
	if err := a.checkPassword(password); err != nil {
		return 0, err
	}
	return a.ID, nil
}

func (uc *AccountUsecase) CheckPasswordByEmail(ctx context.Context, email, password string) (int64, error) {
	a, err := uc.repo.FindByEmail(ctx, email)
	if err != nil {
		return 0, err
	}
	if err := a.checkPassword(password); err != nil {
		return 0, err
	}
	return a.ID, nil
}

func (uc *AccountUsecase) CheckPasswordByID(ctx context.Context, id int64, password string) (int64, error) {
	a, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return 0, err
	}
	if err := a.checkPassword(password); err != nil {
		return 0, err
	}
	return a.ID, nil
}

func (uc *AccountUsecase) ModifyPassword(ctx context.Context, id int64, oldPassword, newPassword string) error {
	a, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if err := a.checkPassword(oldPassword); err != nil {
		return err
	}
	a.Password = newPassword
	if !a.isPasswordValid() {
		return errors.New("invalid password with regex")
	}
	if err := a.encryptPassword(); err != nil {
		return err
	}
	return uc.repo.Update(ctx, a)
}

func (uc *AccountUsecase) checkAccountExist(ctx context.Context, a Account) error {
	eg := errgroup.Group{}
	eg.Go(func() error {
		if a.Email == "" {
			return nil
		}
		_, err := uc.repo.FindByEmail(ctx, a.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return errors.New("db error")
		}
		return errors.New("email exist")
	})
	eg.Go(func() error {
		if a.Mobile == "" {
			return nil
		}
		_, err := uc.repo.FindByMobile(ctx, a.Mobile)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return errors.New("db error")
		}
		return errors.New("mobile exist")
	})
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (a *Account) isPasswordValid(patterns ...string) bool {
	// check with the given pattern
	if len(patterns) > 0 {
		pattern := patterns[0]
		return utils.IsValidWithRegex(pattern, a.Password)
	}
	// check with the default pattern
	return utils.IsValidWithRegex("^[A-Za-z\\d\\S]{8,}", a.Password) // 密码长度至少 8 位
}

func (a *Account) encryptPassword() error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(a.Password),
		bcrypt.DefaultCost,
	)
	a.Password = string(encryptedPassword)
	return err
}

func (a *Account) checkPassword(password string) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	if string(encryptedPassword) != a.Password {
		return errors.New("invalid password")
	}
	return nil
}
