package model

import (
    // "time"
    _ "github.com/jinzhu/gorm"
)

type Organization struct {
	ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT;"`
    Name                string `gorm:"not null; unique; size:255;"`
    Description         string `gorm:"size:1023;"`
    Email               string `gorm:"not null; unique; size:255;"`
    Status              uint8 `gorm:"default: 1;"`
    OwnerID             uint64 `gorm:"index"`
    StreetAdddress      string `gorm:"size:127;"`
    StreetAdddressExtra string `gorm:"size:127;"`
    City                string `gorm:"size:127;"`
    Province            string `gorm:"size:127;"`
    Country             string `gorm:"size:127;"`
    Currency            string  `gorm:"type:varchar(3)"`
    Language            string `gorm:"size:2;"`
    Website             string `gorm:"size:127;"`
    Phone               string `gorm:"size:10;"`
    Fax                 string `gorm:"size:10;"`
}

// Give custom table name in our database.
func (u Organization) TableName() string {
    return "cc_organizations"
}
