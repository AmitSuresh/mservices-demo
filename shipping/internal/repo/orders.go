package repo

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Customer string `gorm:"column:customer_name"`
	Sku      string `gorm:"column:sku;unique"`
	Quantity uint32 `gorm:"column:quantity"`
	Id       string `gorm:"column:id;unique;primaryKey"`
	Shipped  bool   `gorm:"column:shipped"`
	ETA      string `gorm:"column:eta"`
}

type OrderShipping struct {
	Id      string `gorm:"column:id;unique;primaryKey"`
	OrderId string `gorm:"column:order_id"`
	Shipped bool   `gorm:"column:shipped"`
	ETA     string `gorm:"column:eta"`

	Order Order `gorm:"foreignKey:OrderId"`
}

type OrderShippings struct {
	mu            *sync.Mutex
	shippedOrders []*OrderShipping
}

/* func (s *OrderShipping) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&Order{}).
		Where("id = ?", s.OrderId).
		UpdateColumns(map[string]interface{}{
			"Shipped": true,
			"ETA":     s.ETA,
		}).Error
} */

func (s *OrderShipping) BeforeCreate(tx *gorm.DB) (err error) {
	var count int64
	if s.OrderId == "" {
		return errors.New("OrderId cannot be empty")
	}
	err = tx.Model(&Order{}).
		Where("id = ?", s.OrderId).
		Count(&count).
		Error
	if err != nil {
		return err
	}
	var order Order
	err = tx.First(&order, "id = ?", s.OrderId).Error
	if err != nil {
		return err // Return error if Order is not found
	}
	println(order.Customer)
	s.Id = uuid.NewString()

	if count > 0 {
		return tx.Model(&Order{}).
			Where("id = ?", s.Order.Id).
			UpdateColumns(map[string]interface{}{
				"Shipped": true,
				"ETA":     s.ETA,
			}).Error
	} else {
		return errors.New("no records found")
	}

}
