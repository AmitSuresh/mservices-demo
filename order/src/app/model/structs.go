package model

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
