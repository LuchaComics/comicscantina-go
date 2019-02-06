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
type Order struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    User                User `gorm:"foreignkey:UserID"` // Model
    UserID              uint64 `gorm:"index"`
    Organization        Organization `gorm:"foreignkey:OrganizationID"` // Model
    OrganizationID      uint64 `gorm:"index;"`
    Store               Store `gorm:"foreignkey:StoreID"` // Model
    StoreID             uint64 `gorm:"index;"`
    BillingStreetAddress        string `gorm:"size:127;"`
    BillingStreetAddressExtra   string `gorm:"size:127;"`
    BillingCity                 string `gorm:"size:127;"`
    BillingProvince             string `gorm:"size:127;"`
    BillingCountry              string `gorm:"size:127;"`
    BillingPostal               string `gorm:"size:31;"`
    ShippingStreetAddress       string `gorm:"size:127;"`
    ShippingStreetAddressExtra  string `gorm:"size:127;"`
    ShippingCity                string `gorm:"size:127;"`
    ShippingProvince            string `gorm:"size:127;"`
    ShippingCountry             string `gorm:"size:127;"`
    ShippingPostal              string `gorm:"size:31;"`
    OrderDate           time.Time
    ShipDate            time.Time
    RequiredDate        time.Time
    Shipper             Shipper `gorm:"foreignkey:ShipperID"` // Model
    ShipperID           uint64 `gorm:"index"`
    SubAmount           decimal.Decimal `json:"sub_amount" sql:"type:decimal(20,8);"`
    FreightAmount       decimal.Decimal `json:"freight_amount" sql:"type:decimal(20,8);"`
    TaxRate             float64
    TaxAmount           decimal.Decimal `json:"tax_amount" sql:"type:decimal(20,8);"`
    CreatedAt           time.Time
    UpdatedAt           time.Time
    TransactStatus      uint8 `gorm:"DEFAULT: 1;"`
    SubAmountWithTax    decimal.Decimal `json:"sub_amount_with_tax" sql:"type:decimal(20,8);"`
    DiscountAmount      decimal.Decimal `json:"discount_amount" sql:"type:decimal(20,8);"`
    DiscountType        uint8 `gorm:"DEFAULT: 0;"`
    TotalAmount         decimal.Decimal `json:"total_amount" sql:"type:decimal(20,8);"`
    ErrorCode           string `gorm:"size:127;"`
    ErrorMessage        string `gorm:"size:127;"`
}

// Give custom table name in our database.
func (u Order) TableName() string {
    return "cc_orders"
}
