package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
)

//a struct to rep organization.
type Organization struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Name                string `gorm:"not null; unique; size:255;"`
    Description         string `gorm:"size:1023;"`
    Email               string `gorm:"not null; size:255;"`
    Status              uint8 `gorm:"DEFAULT: 1;"`
    OwnerID             uint64 `gorm:"index"`
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
    Facebook            string `gorm:"size:255;"`
    Twitter             string `gorm:"size:255;"`
    YouTube             string `gorm:"size:255;"`
    Google              string `gorm:"size:255;"`
    Employees           []User `gorm:"foreignkey:OrganizationID;"`
    Stores              []Store `gorm:"foreignkey:OrganizationID;"`
    Products            []Product `gorm:"foreignkey:OrganizationID;"`
}

// Status
// 1 = Active
// 2 = Inactive

// Give custom table name in our database.
func (u Organization) TableName() string {
    return "cc_organizations"
}
