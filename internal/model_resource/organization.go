package model_resource

import (
    _ "github.com/jinzhu/gorm"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

// func (dao *database.DataAcessObject) LookupUserByEmail

// Function looks up the user by email.
func DBLookupOrganizationByID(id uint64) (*model.Organization, int) {
    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // The model we will be creating.
    var organization model.Organization
    var count int

    // Find our user.
    db.Where("id = ?", id).First(&organization).Count(&count)
    return &organization, count
}


func DBNewOrganization(name string, description string, email string, ownerID uint64) (*model.Organization, error) {
    // The model we will be creating.
    var organization model.Organization

    // Create our `User` object in our database.
    organization = model.Organization {
        Email:        email,
        Name:         name,
        Description:  description,
        OwnerID:      ownerID,
    }

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our object in the database.
    db.Create(&organization)

    return &organization, nil
}
