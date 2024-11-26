package service

import (
	"context"
	"errors"

	pb "github.com/LXJ0000/tok/backend/basicService/api/basicService/v1"
	"github.com/LXJ0000/tok/backend/basicService/internal/biz"
)

type AccountServiceService struct {
	pb.UnimplementedAccountServiceServer
	accountUc *biz.AccountUsecase
}

func NewAccountServiceService(accountUc *biz.AccountUsecase) *AccountServiceService {
	return &AccountServiceService{accountUc: accountUc}
}

func (s *AccountServiceService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	a := &biz.Account{
		Mobile:   req.Mobile,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := s.accountUc.CreateAccount(ctx, a); err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		Meta:      SuccessMeta,
		AccountId: a.ID,
	}, nil
}

func (s *AccountServiceService) CheckAccount(ctx context.Context, req *pb.CheckAccountRequest) (*pb.CheckAccountResponse, error) {
	var (
		id  int64
		err error
	)
	if req.AccountId != 0 {
		if id, err = s.accountUc.CheckPasswordByID(ctx, req.AccountId, req.Password); err != nil {
			return nil, err
		}
	} else if req.Mobile != "" {
		if id, err = s.accountUc.CheckPasswordByMobile(ctx, req.Mobile, req.Password); err != nil {
			return nil, err
		}
	} else if req.Email != "" {
		if id, err = s.accountUc.CheckPasswordByEmail(ctx, req.Email, req.Password); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid request")
	}
	return &pb.CheckAccountResponse{
		Meta:      SuccessMeta,
		AccountId: id,
	}, nil
}

func (s *AccountServiceService) Bind(ctx context.Context, req *pb.BindRequest) (*pb.BindResponse, error) {
	return &pb.BindResponse{}, nil
}

func (s *AccountServiceService) Unbind(ctx context.Context, req *pb.UnbindRequest) (*pb.UnbindResponse, error) {
	return &pb.UnbindResponse{}, nil
}
