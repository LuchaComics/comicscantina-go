package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
)

//a struct to rep organization's store.
type Store struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT;"`
    Name                string `gorm:"not null; unique; size:255;"`
    Description         string `gorm:"size:1023;"`
    Email               string `gorm:"not null; unique; size:255;"`
    Status              uint8 `gorm:"default: 1;"`
    StreetAddress       string `gorm:"size:127;"`
    StreetAddressExtra  string `gorm:"size:127;"`
    City                string `gorm:"size:127;"`
    Province            string `gorm:"size:127;"`
    Country             string `gorm:"size:127;"`
    Currency            string `gorm:"size:3;"`
    Language            string `gorm:"size:2;"`
    Website             string `gorm:"size:127;"`
    Phone               string `gorm:"size:10;"`
    Fax                 string `gorm:"size:10;"`
    CreatedAt           time.Time
    UpdatedAt           time.Time
    OrganizationID      uint64 `gorm:"index;"`
    Employees           []User `gorm:"foreignkey:EmployeeInStoreID;"`
    Products            []Product `gorm:"foreignkey:StoreID;"`
}

// Give custom table name in our database.
func (u Store) TableName() string {
    return "cc_stores"
}
