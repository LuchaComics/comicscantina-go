package model

import (
    _ "github.com/jinzhu/gorm"
)


type Supplier struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Name                string `gorm:"not null; size:255;"`
    Organization        Organization `gorm:"foreignkey:OrganizationID"` // Model
    OrganizationID      uint64 `gorm:"index;"`
}

// Give custom table name in our database.
func (u Supplier) TableName() string {
    return "cc_suppliers"
}
