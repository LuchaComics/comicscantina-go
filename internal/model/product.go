package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
    "github.com/shopspring/decimal"
)

// Status
// (1) New
// (2) Sold

//a struct to rep product.
type Product struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Name                string `gorm:"not null; size:255;"`
    ShortDescription    string `gorm:"type:varchar(127)"`
    LongDescription     string `gorm:"type:text"`
    Status              uint8 `gorm:"DEFAULT: 1;"`
    OrganizationID      uint64 `gorm:"index;"`
    StoreID             uint64 `gorm:"index;"`
    CreatedAt           time.Time
    UpdatedAt           time.Time
    SKU                 string `gorm:"type:varchar(127)"`
    IDSKU               string `gorm:"type:varchar(127)"`
    VendorProductID     string `gorm:"type:varchar(127)"`
    // Category            Order `gorm:"foreignkey:Category"` // Model
    // SupplierID          uint64 `gorm:"index;"`
    Category            Category `gorm:"foreignkey:Category"` // Model
    CategoryID          uint64 `gorm:"index;"`
    MSRP                decimal.Decimal `json:"unit_price" sql:"type:decimal(20,8);"`
    Price           decimal.Decimal `json:"price" sql:"type:decimal(20,8);"`
    // Picture //TODO
    // Ranking //TODO
}

// Give custom table name in our database.
func (u Product) TableName() string {
    return "cc_products"
}
