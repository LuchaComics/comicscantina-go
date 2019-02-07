package model

import (
    _ "github.com/jinzhu/gorm"
)


type Category struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Name                string `gorm:"not null; size:255;"`
    ShortDescription    string `gorm:"type:varchar(127)"`
    LongDescription     string `gorm:"type:text"`
    Organization        Organization `gorm:"foreignkey:OrganizationID"` // Model
    OrganizationID      uint64 `gorm:"index;"`
}

// Give custom table name in our database.
func (u Category) TableName() string {
    return "cc_categories"
}
