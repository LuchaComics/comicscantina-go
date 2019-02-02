package model_manager

import (
    _ "github.com/jinzhu/gorm"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

/* The structure of our manager. */

type OrganizationManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var organizationManager *OrganizationManager


/* The mangaer functions */

func OrganizationManagerInstance() (*OrganizationManager) {
    if organizationManager != nil {
        return organizationManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        organizationManager = &OrganizationManager{dao}
        return organizationManager
    }
}

func (manager *OrganizationManager) GetByID(id uint64) (*model.Organization, int) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.Organization // The model we will be returning.
    var count int
    orm.Where("id = ?", id).First(&org).Count(&count) // Find our user.
    return &org, count
}

func (manager *OrganizationManager) GetByName(name string) (*model.Organization, int) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.Organization // The model we will be returning.
    var count int
    orm.Where("name = ?", name).First(&org).Count(&count) // Find our user.
    return &org, count
}
