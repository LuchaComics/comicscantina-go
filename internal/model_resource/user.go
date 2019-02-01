package model_resource

import (
    _ "github.com/jinzhu/gorm"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/base/service"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

// func (dao *database.DataAcessObject) LookupUserByEmail

// Function looks up the user by email.
func DBLookupUserByEmail(email string) (*model.User, int) {
    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // The model we will be creating.
    var user model.User
    var count int

    // Find our user.
    db.Where("email = ?", email).First(&user).Count(&count)
    return &user, count
}

// Function looks up the user by id.
func DBLookupUserByID(id uint64) (*model.User, int) {
    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // The model we will be creating.
    var user model.User
    var count int

    // Find our user.
    db.Where("id = ?", id).First(&user).Count(&count)
    return &user, count
}

func DBNewUser(email string, password string, firstName string, lastName string) (*model.User, error) {
    // The model we will be creating.
    var user model.User

    // Secure our password so it's stored in an unreadable form.
    hashedPassword, _ := service.HashPassword(password)

    // Create our `User` object in our database.
    user = model.User {
        Email:        email,
        PasswordHash: hashedPassword,
        FirstName:    firstName,
        LastName:     lastName,
    }

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our object in the database.
    db.Create(&user)

    return &user, nil

}
