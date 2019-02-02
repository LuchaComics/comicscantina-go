package model_manager

import (
    _ "github.com/jinzhu/gorm"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/base/utils"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

/* The structure of our manager. */

type StoreManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var storeManager *StoreManager


/* The mangaer functions */

func StoreManagerInstance() (*StoreManager) {
    if storeManager != nil {
        return storeManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        storeManager = &StoreManager{dao}
        return storeManager
    }
}

func (manager *StoreManager) GetByID(id uint64) (*model.Store, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.Store // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&org).Count(&count) // Find our user.
    return &org, count
}

func (manager *StoreManager) GetByName(name string) (*model.Store, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.Store // The model we will be returning.
    var count uint64
    orm.Where("name = ?", name).First(&org).Count(&count) // Find our user.
    return &org, count
}

func (manager *StoreManager) AllByPageIndex(pageIndex uint64) ([]model.Store, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.

    // Initial array to hold query results.
    var stores []model.Store

    // Where will find all records.
    orm.Where("id > ?", 0)

    // Make our paginated query.
    pagination := utils.Pagging(&utils.Param{
		DB:      orm,
		Page:    pageIndex,
		Limit:   25,
		OrderBy: []string{"id asc"},
		ShowSQL: false,
	}, &stores)

    return stores, pagination.TotalRecord
}
