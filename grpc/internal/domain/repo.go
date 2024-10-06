package domain

import (
	"github.com/AmitSuresh/grpc/proto/order_service"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type orderRepo struct {
	l  *zap.Logger
	db *gorm.DB
}

type OrderRepo interface {
	CreateOrder(*order_service.OrderReq) (*Order, error)
}

func (repo *orderRepo) CreateOrder(req *order_service.OrderReq) (*Order, error) {
	o := &Order{
		Customer: req.GetCustomer(),
		Sku:      req.GetSku(),
		Quantity: req.GetQuantity(),
		Id:       uuid.New().String(),
		Shipped:  false,
	}
	tx := repo.db.Begin()
	if tx.Error != nil {
		repo.l.Error("error starting tx", zap.Error(tx.Error))
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			repo.l.Error("error starting tx", zap.Any("r", r))
			tx.Rollback()
		}
	}()

	if err := tx.Create(&o).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return o, nil
}

func NewOrderRepo(l *zap.Logger, db *gorm.DB) OrderRepo {
	return &orderRepo{
		l:  l,
		db: db,
	}
}
