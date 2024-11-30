package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Minio struct {
	core *minio.Core
}

func NewMinio(core *minio.Core) *Minio {
	return &Minio{core: core}
}

// PreSignGetUrl returns a pre-signed URL for downloading an object
func (r *Minio) PreSignGetUrl(ctx context.Context, bucketName, objectName, fileName string, expire time.Duration) (string, error) {
	reqParams := make(url.Values)
	if fileName != "" {
		reqParams.Set("response-content-disposition", "attachment; filename="+fileName)
	}
	signedUrl, err := r.core.PresignedGetObject(ctx, bucketName, objectName, expire, reqParams)
	if err != nil {
		return "", err
	}
	return signedUrl.String(), nil
}

// PreSignPutUrl returns a pre-signed URL for uploading an object
func (r *Minio) PreSignPutUrl(ctx context.Context, bucketName, objectName string, expire time.Duration) (string, error) {
	signedUrl, err := r.core.PresignedPutObject(ctx, bucketName, objectName, expire)
	if err != nil {
		return "", err
	}
	return signedUrl.String(), nil
}

// CreateSlicingUpload creates a new multipart upload
func (r *Minio) CreateSlicingUpload(ctx context.Context, bucketName, objectName string, options minio.PutObjectOptions) (uploadId string, err error) {
	return r.core.NewMultipartUpload(ctx, bucketName, objectName, options)
}

// ListSlicingFileParts lists the parts of a multipart upload
func (r *Minio) ListSlicingFileParts(ctx context.Context, bucketName, objectName, uploadId string, partsNum int64) (minio.ListObjectPartsResult, error) {
	var nextPartNumberMarker int
	return r.core.ListObjectParts(ctx, bucketName, objectName, uploadId, nextPartNumberMarker, int(partsNum)+1)
}

// PreSignSlicingPutUrl returns a pre-signed URL for uploading a part of a multipart upload
func (r *Minio) PreSignSlicingPutUrl(ctx context.Context, bucketName, objectName, uploadId string, parts int64) (string, error) {
	params := url.Values{
		"uploadId":   {uploadId},
		"partNumber": {strconv.FormatInt(parts, 10)},
	}
	signedUrl, err := r.core.Presign(ctx, http.MethodPut, bucketName, objectName, time.Hour, params)
	if err != nil {
		return "", err
	}
	return signedUrl.String(), nil
}

// MergeSlices merges the parts of a multipart upload
func (r *Minio) MergeSlices(ctx context.Context, bucketName, objectName, uploadId string, parts []minio.CompletePart) error {
	_, err := r.core.CompleteMultipartUpload(ctx, bucketName, objectName, uploadId, parts, minio.PutObjectOptions{})
	return err
}

// GetObjectHash returns the hash of an object
func (r *Minio) GetObjectHash(ctx context.Context, bucketName, objectName string) (string, error) {
	stats, err := r.core.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return "", err
	}
	return strings.ToUpper(stats.ETag), nil
}
