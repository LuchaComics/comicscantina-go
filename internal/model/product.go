package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
)

// Status
// (1) New
// (2) Sold

//a struct to rep product.
type Product struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Name                string `gorm:"not null; size:255;"`
    Status              uint8 `gorm:"DEFAULT: 1;"`
    CreatedAt           time.Time
    UpdatedAt           time.Time
    OrganizationID      uint64 `gorm:"index;"`
    StoreID             uint64 `gorm:"index;"`
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
func (u Product) TableName() string {
    return "cc_products"
}
