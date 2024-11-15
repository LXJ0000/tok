package service

import (
	"context"

	pb "github.com/LXJ0000/tok/backend/basicService/api/basicService/v1"
)

type PostServiceService struct {
	pb.UnimplementedPostServiceServer
}

func NewPostServiceService() *PostServiceService {
	return &PostServiceService{}
}

func (s *PostServiceService) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateResponse, error) {
	return &pb.CreateTemplateResponse{}, nil
}
func (s *PostServiceService) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest) (*pb.UpdateTemplateResponse, error) {
	return &pb.UpdateTemplateResponse{}, nil
}
func (s *PostServiceService) ListTemplate(ctx context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateResponse, error) {
	return &pb.ListTemplateResponse{}, nil
}
func (s *PostServiceService) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
	return &pb.GetTemplateResponse{}, nil
}
func (s *PostServiceService) RemoveTemplate(ctx context.Context, req *pb.RemoveTemplateRequest) (*pb.RemoveTemplateResponse, error) {
	return &pb.RemoveTemplateResponse{}, nil
}
func (s *PostServiceService) SendSms(ctx context.Context, req *pb.SendSmsRequest) (*pb.SendSmsResponse, error) {
	return &pb.SendSmsResponse{}, nil
}
func (s *PostServiceService) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	return &pb.SendEmailResponse{}, nil
}
func (s *PostServiceService) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	return &pb.SendResponse{}, nil
}
