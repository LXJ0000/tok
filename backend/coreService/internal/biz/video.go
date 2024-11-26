package biz

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/LXJ0000/tok/backend/coreService/internal/infra/utils"
	"github.com/go-kratos/kratos/v2/log"
)

// Video is a Video model.
type Video struct {
	ID           int64     `json:"id"`            // 视频ID
	UserID       int64     `json:"user_id"`       // 用户ID
	Title        string    `json:"title"`         // 标题
	Description  string    `json:"description"`   // 描述
	VideoURL     string    `json:"video_url"`     // 视频URL
	CoverURL     string    `json:"cover_url"`     // 封面URL
	LikeCount    int64     `json:"like_count"`    // 点赞数
	CollectCount int64     `json:"collect_count"` // 收藏数
	ForwardCount int64     `json:"forward_count"` // 转发数
	CommentCount int64     `json:"comment_count"` // 评论数
	IsDeleted    bool      `json:"is_deleted"`    // 是否删除
	CreateTime   time.Time `json:"create_time"`   // 创建时间
	UpdateTime   time.Time `json:"update_time"`   // 更新时间

	// Extra
	Author     *User `json:"author"`      // 作者信息
	IsFavorite int64 `json:"is_favorite"` // 是否点赞
	IsCollect  int64 `json:"is_collect"`  // 是否收藏
}

// VideoRepo is a Greater repo.
type VideoRepo interface {
	Save(context.Context, *Video) error
	Update(context.Context, *Video) error
	FindByID(context.Context, int64) (*Video, error)
	FindByIDList(context.Context, []int64) ([]*Video, error)
	FindByUserID(ctx context.Context, userID int64, lastID int64, limit int) ([]*Video, error)
	Find(ctx context.Context, lastID int64, limit int) ([]*Video, error)
}

// VideoUsecase is a Video usecase.
type VideoUsecase struct {
	repo     VideoRepo
	userRepo UserRepo
	log      *log.Helper
}

// NewVideoUsecase new a Video usecase.
func NewVideoUsecase(repo VideoRepo, userRepo UserRepo, logger log.Logger) *VideoUsecase {
	return &VideoUsecase{repo: repo, userRepo: userRepo, log: log.NewHelper(logger)}
}

// CreateVideo creates a Video, and returns the new Video.
func (uc *VideoUsecase) CreateVideo(ctx context.Context, g *Video) error {
	g.ID = utils.GenID()
	return uc.repo.Save(ctx, g)
}

// UpdateVideo updates the Video, and returns the updated Video.
func (uc *VideoUsecase) UpdateVideo(ctx context.Context, g *Video) error {
	return uc.repo.Update(ctx, g)
}

// GetVideo gets the Video by the ID.
func (uc *VideoUsecase) GetVideo(ctx context.Context, id int64) (*Video, error) {
	v, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	author, err := uc.userRepo.FindByID(ctx, v.UserID)
	if err != nil {
		return nil, err
	}
	v.Author = author
	return v, nil
}

func (uc *VideoUsecase) GetVideoByIdList(ctx context.Context, videoIdList []int64) ([]*Video, error) {
	// 目前用不到 优化点：查询用户信息考虑去重
	vs, err := uc.repo.FindByIDList(ctx, videoIdList)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	wg.Add(len(vs))
	for i := range vs {
		go func(i int) {
			defer wg.Done()
			author, err := uc.userRepo.FindByID(ctx, vs[i].UserID)
			if err != nil {
				log.Error("find user by id failed", "error", err, "video", vs[i])
				return
			}
			vs[i].Author = author
		}(i)
	}
	wg.Wait()
	return vs, nil
}

func (uc *VideoUsecase) Feed(ctx context.Context, userID, lastTime int64, limit int) ([]*Video, error) {
	if lastTime <= 0 {
		lastTime = math.MaxInt64
	}
	vs, err := uc.repo.Find(ctx, lastTime, limit)
	if err != nil {
		return nil, err
	}
	users := make(map[int64]*User)
	userIDs := make([]int64, 0, len(users))
	for _, v := range vs {
		users[v.UserID] = nil
	}
	for userID := range users {
		userIDs = append(userIDs, userID)
	}
	us, err := uc.userRepo.FindByIDList(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	for _, u := range us {
		users[u.ID] = u
	}
	for _, v := range vs {
		v.Author = users[v.UserID]
	}
	return vs, nil
}
