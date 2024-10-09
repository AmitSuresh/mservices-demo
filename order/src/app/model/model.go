package model

import (
	"sync"

	"go.uber.org/zap"
)

type CreateOrder struct {
	CustomerName string `json:"customer_name"`
	Quantity     uint32 `json:"quantity"`
	Sku          string `json:"sku"`
}

type OrderShipping struct {
	Id      string `gorm:"column:id;unique;primaryKey"`
	OrderId string `gorm:"column:order_id"`
	Shipped bool   `gorm:"column:shipped"`
	ETA     string `gorm:"column:eta"`

	//Order Order `gorm:"foreignKey:OrderId"`
}
type ShipmentDetails struct {
	Mu     sync.Mutex
	Orders []OrderShipping
}

/*
type Order struct {
	Customer string `gorm:"column:customer_name"`
	Sku      string `gorm:"column:sku;unique"`
	Quantity uint32 `gorm:"column:quantity"`
	Id       string `gorm:"column:id;unique;primaryKey"`
	Shipped  bool   `gorm:"column:shipped"`
	ETA      string `gorm:"column:eta"`
}
*/

func ProcessOrderShipping(o *OrderShipping, so *ShipmentDetails, l *zap.Logger) {
	so.Mu.Lock()
	so.Orders = append(so.Orders, *o)
	l.Info("updated Shipping list", zap.Any("Orders", so.Orders))
	so.Mu.Unlock()
}
