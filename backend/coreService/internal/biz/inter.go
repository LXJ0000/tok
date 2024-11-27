package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Inter is a Inter model.
type Inter struct {
	ID         int64     `json:"id"`
	BizID      int64     `json:"biz_id"`
	Biz        string    `json:"biz"`
	ReadCnt    int64     `json:"read_cnt"`
	LikeCnt    int64     `json:"like_cnt"`
	CollectCnt int64     `json:"collect_cnt"`
	IsDeleted  int32     `json:"is_deleted"`  // 是否删除
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

// InterRepo is a Greater repo.
type InterRepo interface {
	IncrReadCount(c context.Context, biz string, id int64) error
	Like(c context.Context, biz string, bizID, userID int64) error
	CancelLike(c context.Context, biz string, bizID, userID int64) error
	Collect(c context.Context, biz string, bizID, userID, collectionID int64) error
	CancelCollect(c context.Context, biz string, bizID, userID, collectionID int64) error
	Stat(c context.Context, biz string, bizID, userID int64) (*InterStat, error)
}

// InterUsecase is a Inter usecase.
type InterUsecase struct {
	repo InterRepo
	log  *log.Helper
}

// NewInterUsecase new a Inter usecase.
func NewInterUsecase(repo InterRepo, logger log.Logger) *InterUsecase {
	return &InterUsecase{repo: repo, log: log.NewHelper(logger)}
}

type InterStat struct {
	LikeCnt    int64 `json:"like_count"`
	CollectCnt int64 `json:"collect_count"`
	ReadCnt    int64 `json:"read_count"`
	IsLike     bool  `json:"is_like"`
	IsCollect  bool  `json:"is_collect"`
}

// IncrReadCount .
func (uc *InterUsecase) IncrReadCount(ctx context.Context, biz string, id int64) error {
	return uc.repo.IncrReadCount(ctx, biz, id)
}

// Like .
func (uc *InterUsecase) Like(ctx context.Context, biz string, bizID, userID int64) error {
	return uc.repo.Like(ctx, biz, bizID, userID)
}

// CancelLike .
func (uc *InterUsecase) CancelLike(ctx context.Context, biz string, bizID, userID int64) error {
	return uc.repo.CancelLike(ctx, biz, bizID, userID)
}

// Collect .
func (uc *InterUsecase) Collect(ctx context.Context, biz string, bizID, userID, collectionID int64) error {
	return uc.repo.Collect(ctx, biz, bizID, userID, collectionID)
}

// CancelCollect .
func (uc *InterUsecase) CancelCollect(ctx context.Context, biz string, bizID, userID, collectionID int64) error {
	return uc.repo.CancelCollect(ctx, biz, bizID, userID, collectionID)
}

// Stat .
func (uc *InterUsecase) Stat(ctx context.Context, biz string, bizID, userID int64) (*InterStat, error) {
	return uc.repo.Stat(ctx, biz, bizID, userID)
}
