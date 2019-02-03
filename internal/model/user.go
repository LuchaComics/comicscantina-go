package model

import (
    _ "github.com/jinzhu/gorm"
)

//a struct to rep user account
type User struct {
    ID                uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Email             string `gorm:"not null; unique; size:255;"`
    PasswordHash      string `gorm:"size:511`
    FirstName         string `gorm:"type:varchar(127)”`
    LastName          string `gorm:"type:varchar(127)”`
    EmployerID        uint64 `gorm:"index;"` // ID of organization this user is an employee of.
    OrganizationID    uint64 `gorm:"index;"` // ID of organization this user owns
    Receipts          []Receipt `gorm:"foreignkey:OwnerID"`
    Status            uint8 `gorm:"DEFAULT: 1;"` // 1 = Active, 2 = Inactive
    GroupID           uint8 `gorm:"DEFAULT: 1;"` // 1 = Regular, 2 = Admin
}

// Give custom table name in our database.
func (u User) TableName() string {
    return "cc_users"
}
