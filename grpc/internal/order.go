package internal

import (
	"context"
	"errors"

	"github.com/AmitSuresh/grpc/internal/domain"
	"github.com/AmitSuresh/grpc/proto/order_service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type orderService struct {
	repo domain.OrderRepo
	l    *zap.Logger
}

type OrderService interface {
	CreateOrder(context.Context, *order_service.OrderReq) (*order_service.OrderResp, error)
}

func (ser *orderService) CreateOrder(ctx context.Context, req *order_service.OrderReq) (*order_service.OrderResp, error) {
	o, err := ser.repo.CreateOrder(req)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case errors.Is(err, gorm.ErrInvalidData):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}
	resp := &order_service.OrderResp{
		Customer: o.Customer,
		Id:       o.Id,
		Shipped:  o.Shipped,
	}
	return resp, nil
}

func NewOrderService(repo domain.OrderRepo, l *zap.Logger) order_service.OrderServiceServer {
	return &orderService{
		repo: repo,
		l:    l,
	}
}
