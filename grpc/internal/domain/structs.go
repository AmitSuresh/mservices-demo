package domain

type Order struct {
	Customer string `gorm:"column:customer_name"`
	Sku      string `gorm:"column:sku;unique"`
	Quantity uint32 `gorm:"column:quantity"`
	Id       string `gorm:"column:id;unique;primaryKey"`
	Shipped  bool   `gorm:"column:shipped"`
}
