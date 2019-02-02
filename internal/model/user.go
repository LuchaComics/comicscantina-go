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
    Organizations     []Organization `gorm:"foreignkey:OwnerID"`
    EmployeeInStoreID uint64 `gorm:"index"`
}

// Give custom table name in our database.
func (u User) TableName() string {
    return "cc_users"
}
