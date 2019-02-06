package model

import (
    _ "github.com/jinzhu/gorm"
    "github.com/shopspring/decimal"
)


type OrderItem struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Order               Order `gorm:"foreignkey:OrderID"` // Model
    OrderID             uint64 `gorm:"index"`
    Product             Product `gorm:"foreignkey:Product"` // Model
    ProductID           uint64 `gorm:"index"`
    Price               decimal.Decimal `json:"price" sql:"type:decimal(20,8);"`
}

// Give custom table name in our database.
func (u OrderItem) TableName() string {
    return "cc_order_items"
}
