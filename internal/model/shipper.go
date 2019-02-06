package model

import (
    _ "github.com/jinzhu/gorm"
)


type Shipper struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Name                string `gorm:"not null; size:255;"`
    Email               string `gorm:"size:255;"`
    Phone               string `gorm:"size:10;"`
}

// Give custom table name in our database.
func (u Shipper) TableName() string {
    return "cc_shippers"
}
