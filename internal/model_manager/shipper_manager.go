package model_manager

import (
    _ "github.com/jinzhu/gorm"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/base/utils"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

/* The structure of our manager. */

type ShipperManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var shipperManager *ShipperManager


/* The mangaer functions */

func ShipperManagerInstance() (*ShipperManager) {
    if shipperManager != nil {
        return shipperManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        shipperManager = &ShipperManager{dao}
        return shipperManager
    }
}

func (manager *ShipperManager) GetByID(id uint64) (*model.Shipper, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var shipper model.Shipper // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&shipper).Count(&count) // Find our shipper.
    return &shipper, count
}

func (manager *ShipperManager) PaginatedAll(pageIndex uint64) ([]model.Shipper, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.

    // Initial array to hold query results.
    var products []model.Shipper

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

func (manager *ShipperManager) Create(name string, orgID uint64) (*model.Shipper, error) {
    // The model we will be creating.
    var shipper model.Shipper

    // Create our `Shipper` object in our database.
    shipper = model.Shipper {
        Name:             name,
        OrganizationID: orgID,
    }

    orm := manager.dao.GetORM() // Get our database layer.
    orm.Create(&shipper) // Create our object in the database.
    return &shipper, nil
}
