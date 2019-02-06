package model_manager

import (
    _ "github.com/jinzhu/gorm"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/base/utils"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

/* The structure of our manager. */

type ProductManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var productManager *ProductManager


/* The mangaer functions */

func ProductManagerInstance() (*ProductManager) {
    if productManager != nil {
        return productManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        productManager = &ProductManager{dao}
        return productManager
    }
}

func (manager *ProductManager) GetByID(id uint64) (*model.Product, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.Product // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&org).Count(&count) // Find our user.
    return &org, count
}

func (manager *ProductManager) GetByName(name string) (*model.Product, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.Product // The model we will be returning.
    var count uint64
    orm.Where("name = ?", name).First(&org).Count(&count) // Find our user.
    return &org, count
}

func (manager *ProductManager) PaginatedAll(pageIndex uint64) ([]model.Product, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.

    // Initial array to hold query results.
    var products []model.Product

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

func (manager *ProductManager) PaginatedFilterBy(orgID uint64, storeID uint64, pageIndex uint64) ([]model.Product, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.

    // Initial array to hold query results.
    var products []model.Product

    // Filter by `organization_id`.
    if orgID > 0 {
        orm.Where("organization_id = ?", orgID)
    }

    // Filter by `store_id`.
    if orgID > 0 {
        orm.Where("store_id = ?", storeID)
    }

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
