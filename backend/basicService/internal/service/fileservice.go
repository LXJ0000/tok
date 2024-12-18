package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/basicService/api/basicService/v1"
	"github.com/LXJ0000/tok/backend/basicService/internal/biz"
)

type FileServiceService struct {
	pb.UnimplementedFileServiceServer

	fileUc *biz.FileUsecase
}

func NewFileServiceService(fileUc *biz.FileUsecase) *FileServiceService {
	return &FileServiceService{fileUc: fileUc}
}

func (s *FileServiceService) PreSignGet(ctx context.Context, req *pb.PreSignGetRequest) (*pb.PreSignGetResponse, error) {
	g := &biz.File{
		ID:         req.FileContext.FileId,
		DomainName: req.FileContext.Domain,
		BizName:    req.FileContext.BizName,
		Hash:       req.FileContext.Hash,
		FileType:   req.FileContext.FileType,
		FileSize:   req.FileContext.Size,
		ExpireSeconds: req.FileContext.ExpireSeconds,
		FileName: req.FileContext.Filename,
	}
	s.fileUc.PreSignGet(ctx, g)
	return &pb.PreSignGetResponse{}, nil
}

func (s *FileServiceService) PreSignPut(ctx context.Context, req *pb.PreSignPutRequest) (*pb.PreSignPutResponse, error) {
	return &pb.PreSignPutResponse{}, nil
}

func (s *FileServiceService) ReportUploaded(ctx context.Context, req *pb.ReportUploadedRequest) (*pb.ReportUploadedResponse, error) {
	return &pb.ReportUploadedResponse{}, nil
}

func (s *FileServiceService) PreSignSlicingPut(ctx context.Context, req *pb.PreSignSlicingPutRequest) (*pb.PreSignSlicingPutResponse, error) {
	return &pb.PreSignSlicingPutResponse{}, nil
}

func (s *FileServiceService) GetProgressRate4SlicingPut(ctx context.Context, req *pb.GetProgressRate4SlicingPutRequest) (*pb.GetProgressRate4SlicingPutResponse, error) {
	return &pb.GetProgressRate4SlicingPutResponse{}, nil
}

func (s *FileServiceService) MergeFileParts(ctx context.Context, req *pb.MergeFilePartsRequest) (*pb.MergeFilePartsResponse, error) {
	return &pb.MergeFilePartsResponse{}, nil
}

func (s *FileServiceService) RemoveFile(ctx context.Context, req *pb.RemoveFileRequest) (*pb.RemoveFileResponse, error) {
	return &pb.RemoveFileResponse{}, nil
}

func (s *FileServiceService) GetFileInfoById(ctx context.Context, req *pb.GetFileInfoByIdRequest) (*pb.GetFileInfoByIdResponse, error) {
	return &pb.GetFileInfoByIdResponse{}, nil
}
