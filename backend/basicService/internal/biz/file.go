package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/LXJ0000/tok/backend/basicService/internal/infra/pkg/object_storage/minio"
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

func (f *File) objectName() string {
	return fmt.Sprintf("%s/%d", f.BizName, f.ID)
}

type SlicingFile struct {
	File       *File
	TotalParts int64
	UploadId   string
	UploadUrl  []string
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
	repo  FileRepo
	minio *minio.Minio
	log   *log.Helper
}

// NewFileUsecase new a File usecase.
func NewFileUsecase(repo FileRepo, minio *minio.Minio, logger log.Logger) *FileUsecase {
	return &FileUsecase{
		repo:  repo,
		minio: minio,
		log:   log.NewHelper(logger),
	}
}

// CreateFile creates a File, and returns the new File.
func (uc *FileUsecase) CreateFile(ctx context.Context, g *File) error {
	return uc.repo.Save(ctx, g)
}

// PreSignGet returns the pre-signed URL for downloading the file
func (uc *FileUsecase) PreSignGet(ctx context.Context, g *File) (string, error) {
	uc.repo.FindUploadedByID(ctx, g.ID)
	// return uc.minio.PreSignGetUrl(ctx, g.DomainName, "", g.), nil
	return "", nil
}

// PreSignPut returns the pre-signed URL for uploading the file, and the file ID if the file already exists
func (uc *FileUsecase) PreSignPut(ctx context.Context, g *File) (string, int64, error) {
	return "", 0, nil
}

// CheckFileExistedAndGetFile checks if the file exists and returns the file ID if it does
func (uc *FileUsecase) CheckFileExistedAndGetFile(ctx context.Context, g *File) (int64, bool, error) {
	return 0, false, nil
}

// ReportUploaded reports that the file has been uploaded
func (uc *FileUsecase) ReportUploaded(ctx context.Context, g *File) error {
	return nil
}

// PreSignSlicingPut returns the pre-signed URL for uploading the file in slices
func (uc *FileUsecase) PreSignSlicingPut(ctx context.Context, g *File) (*SlicingFile, error) {
	return nil, nil
}

// GetProgressRate4SlicingPut returns the progress rate of the file upload in slices
func (uc *FileUsecase) GetProgressRate4SlicingPut(ctx context.Context, uploadId string, g *File) (map[string]bool, error) {
	return nil, nil
}

// MergeFileParts merges the file parts uploaded in slices
func (uc *FileUsecase) MergeFileParts(ctx context.Context, uploadId string, g *File) error {
	return nil
}

// RemoveFile removes the file
func (uc *FileUsecase) RemoveFile(ctx context.Context, g *File) error {
	return nil
}

func (uc *FileUsecase) GetInfoById(ctx context.Context, domainName, bizName string, fileId int64) (*File, error) {
	return nil, nil
}
