package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
)

// Status
// (1) New
// (2) Sold

//a struct to rep product.
type Receipt struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT;"`
    Status              uint8 `gorm:"DEFAULT: 1;"`
    CreatedAt           time.Time
    UpdatedAt           time.Time
    OrganizationID      uint64 `gorm:"index;"`
    StoreID             uint64 `gorm:"index;"`
    OwnerID             uint64 `gorm:"index"`
    // sub_price
    // has_tax
    // tax_rate
    // tax_amount
    // sub_price_with_tax
    // discount
    // discount_type
    // price
    // cost
    // currency
    // language
}

// Give custom table name in our database.
func (u Receipt) TableName() string {
    return "cc_receipts"
}
