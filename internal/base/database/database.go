package database

import (
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
        panic("Failed to connect database")
    }
    // defer db.Close() // Handle this in `main.go` so do not uncomment this!

    // PLEASE READ FOR MORE INFORAMTION:
    // http://doc.gorm.io/

    // // Automatically delete previous database schema.
    // db.Debug().DropTableIfExists(&model.User{})
    db.Debug().DropTableIfExists(&model.Organization{})

    // Automatically migrate our database schema.
    db.Debug().AutoMigrate(&model.User{})
    db.Debug().AutoMigrate(&model.Organization{})

    // Keep an instance of our new object.
    dao = &DataAcessObject{
        dbPool: db,
    }

    //Return our database connector.
    return dao
}

func (instance *DataAcessObject) DropAndCreateDatabase() {
    // Automatically delete previous database schema.
    instance.dbPool.Debug().DropTableIfExists(&model.Organization{})
    instance.dbPool.Debug().DropTableIfExists(&model.User{})

    // Automatically migrate our database schema.
    instance.dbPool.Debug().AutoMigrate(&model.User{})
    instance.dbPool.Debug().AutoMigrate(&model.Organization{})
}

func (instance *DataAcessObject) GetORM() (*gorm.DB) {
    return instance.dbPool
}
