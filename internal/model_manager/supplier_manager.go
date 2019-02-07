package model_manager

import (
    _ "github.com/jinzhu/gorm"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/base/utils"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

/* The structure of our manager. */

type SupplierManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var supplierManager *SupplierManager


/* The mangaer functions */

func SupplierManagerInstance() (*SupplierManager) {
    if supplierManager != nil {
        return supplierManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        supplierManager = &SupplierManager{dao}
        return supplierManager
    }
}

func (manager *SupplierManager) GetByID(id uint64) (*model.Supplier, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var supplier model.Supplier // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&supplier).Count(&count) // Find our supplier.
    return &supplier, count
}

func (manager *SupplierManager) PaginatedAll(pageIndex uint64) ([]model.Supplier, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.

    // Initial array to hold query results.
    var products []model.Supplier

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

func (manager *SupplierManager) Create(name string, orgID uint64) (*model.Supplier, error) {
    // The model we will be creating.
    var supplier model.Supplier

    // Create our `Supplier` object in our database.
    supplier = model.Supplier {
        Name:             name,
        OrganizationID: orgID,
    }

    orm := manager.dao.GetORM() // Get our database layer.
    orm.Create(&supplier) // Create our object in the database.
    return &supplier, nil
}
