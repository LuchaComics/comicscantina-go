package model_manager

import (
    _ "github.com/jinzhu/gorm"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/base/utils"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

/* The structure of our manager. */

type CategoryManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var categoryManager *CategoryManager


/* The mangaer functions */

func CategoryManagerInstance() (*CategoryManager) {
    if categoryManager != nil {
        return categoryManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        categoryManager = &CategoryManager{dao}
        return categoryManager
    }
}

func (manager *CategoryManager) GetByID(id uint64) (*model.Category, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var category model.Category // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&category).Count(&count) // Find our category.
    return &category, count
}

func (manager *CategoryManager) PaginatedAll(pageIndex uint64) ([]model.Category, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.

    // Initial array to hold query results.
    var products []model.Category

    // Where will find all records.
    orm.Where("id > ?", 0)

    // Make our paginated query.
    pagination := utils.Pagging(&utils.Param{
		DB:      orm,
		Page:    pageIndex,
		Limit:   25,
		OrderBy: []string{"id asc"},
		ShowSQL: false,
	}, &products)

    return products, pagination.TotalRecord
}

func (manager *CategoryManager) Create(name string, shortDescription string, longDescription string, orgID uint64) (*model.Category, error) {
    // The model we will be creating.
    var category model.Category

    // Create our `Category` object in our database.
    category = model.Category {
        Name:             name,
        ShortDescription: shortDescription,
        LongDescription:  longDescription,
        OrganizationID: orgID,
    }

    orm := manager.dao.GetORM() // Get our database layer.
    orm.Create(&category) // Create our object in the database.
    return &category, nil
}
