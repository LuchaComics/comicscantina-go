package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "github.com/luchacomics/comicscantina-go/internal/base/config"
    "github.com/luchacomics/comicscantina-go/internal/model"
)

/* Database Structure */

type DataAcessObject struct {
    dbPool *gorm.DB
}

/* Global variable. */

var dao *DataAcessObject

/* Private initializer */

func init() {
    Instance()
}

/* Function declaration */

// Function will return an instance of our database access layer (DAO) or the
// function will lazily load the DAO and then return the DAO.
func Instance() (*DataAcessObject) {
    // Lazily load the database connection if it was not created before.
    if dao != nil {
        return dao
    }

    // Get the database configuration text from the environment variables.
    databaseConfigString := config.GetSettingsVariableDatabaseURL()

    // The following code will connect our application to the `postgres` database.
    db, err := gorm.Open("postgres", databaseConfigString)
    if err != nil {
        fmt.Println(err)
        panic("Failed to connect database")
    }
    // defer db.Close() // Handle this in `main.go` so do not uncomment this!

    // PLEASE READ FOR MORE INFORAMTION:
    // http://doc.gorm.io/

    // // Automatically delete previous database schema.
    // db.Debug().DropTableIfExists(&model.Supplier{})
    // db.Debug().DropTableIfExists(&model.Shipper{})
    // db.Debug().DropTableIfExists(&model.Category{})
    // db.Debug().DropTableIfExists(&model.OrderItem{})
    // db.Debug().DropTableIfExists(&model.Order{})
    // db.Debug().DropTableIfExists(&model.Product{})
    // db.Debug().DropTableIfExists(&model.Store{})
    // db.Debug().DropTableIfExists(&model.Organization{})
    // db.Debug().DropTableIfExists(&model.User{})

    // Automatically migrate our database schema.
    db.Debug().AutoMigrate(&model.User{})
    db.Debug().AutoMigrate(&model.Organization{})
    db.Debug().AutoMigrate(&model.Store{})
    db.Debug().AutoMigrate(&model.Product{})
    db.Debug().AutoMigrate(&model.Order{})
    db.Debug().AutoMigrate(&model.OrderItem{})
    db.Debug().AutoMigrate(&model.Category{})
    db.Debug().AutoMigrate(&model.Shipper{})
    db.Debug().AutoMigrate(&model.Supplier{})

    // Keep an instance of our new object.
    dao = &DataAcessObject{
        dbPool: db,
    }

    //Return our database connector.
    return dao
}

func (instance *DataAcessObject) DropAndCreateDatabase() {
    // Automatically delete previous database schema.
    instance.dbPool.Debug().DropTableIfExists(&model.Supplier{})
    instance.dbPool.Debug().DropTableIfExists(&model.Shipper{})
    instance.dbPool.Debug().DropTableIfExists(&model.Category{})
    instance.dbPool.Debug().DropTableIfExists(&model.OrderItem{})
    instance.dbPool.Debug().DropTableIfExists(&model.Order{})
    instance.dbPool.Debug().DropTableIfExists(&model.Product{})
    instance.dbPool.Debug().DropTableIfExists(&model.Store{})
    instance.dbPool.Debug().DropTableIfExists(&model.Organization{})
    instance.dbPool.Debug().DropTableIfExists(&model.User{})

    // Automatically migrate our database schema.
    instance.dbPool.Debug().AutoMigrate(&model.User{})
    instance.dbPool.Debug().AutoMigrate(&model.Organization{})
    instance.dbPool.Debug().AutoMigrate(&model.Store{})
    instance.dbPool.Debug().AutoMigrate(&model.Product{})
    instance.dbPool.Debug().AutoMigrate(&model.Order{})
    instance.dbPool.Debug().AutoMigrate(&model.OrderItem{})
    instance.dbPool.Debug().AutoMigrate(&model.Category{})
    instance.dbPool.Debug().AutoMigrate(&model.Shipper{})
    instance.dbPool.Debug().AutoMigrate(&model.Supplier{})
}

func (instance *DataAcessObject) GetORM() (*gorm.DB) {
    return instance.dbPool
}
