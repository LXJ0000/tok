package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// File is a File model.
type File struct {
	ID         int64     `json:"id"`          // 文件ID
	DomainName string    `json:"domain_name"` // 域名
	BizName    string    `json:"biz_name"`    // 业务名称
	Hash       string    `json:"hash"`        // 文件哈希
	FileSize   int64     `json:"file_size"`   // 文件大小
	FileType   string    `json:"file_type"`   // 文件类型
	Uploaded   int32     `json:"uploaded"`    // 是否上传
	IsDeleted  int32     `json:"is_deleted"`  // 是否删除
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间

}

// FileRepo is a Greater repo.
type FileRepo interface {
	Save(c context.Context, file *File) error
	Update(c context.Context, file *File) error
	FindByID(c context.Context, id int64) (*File, error)
	FindUploadedByID(c context.Context, id int64) (*File, error)
	FindUploadedByHash(c context.Context, hash string) (*File, error)
	DeleteByID(c context.Context, id int64) error
}

// FileUsecase is a File usecase.
type FileUsecase struct {
	repo FileRepo
	log  *log.Helper
}

// NewFileUsecase new a File usecase.
func NewFileUsecase(repo FileRepo, logger log.Logger) *FileUsecase {
	return &FileUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateFile creates a File, and returns the new File.
func (uc *FileUsecase) CreateFile(ctx context.Context, g *File) error {
	return uc.repo.Save(ctx, g)
}
