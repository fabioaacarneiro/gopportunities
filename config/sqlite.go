package config

import (
	"os"

	"github.com/fabioaacarneiro/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
    logger := GetLogger("sqlite")

    dbPath := "./db/main.db"

    // check if the database file exists
    _, err := os.Stat(dbPath)
    if os.IsNotExist(err) {
        logger.Info("database file not found, creating...")

        // create the database file and directory
        err := os.MkdirAll("./db", os.ModePerm)
        if err != nil {
            return nil, err
        }

        file, err := os.Create(dbPath)

        if err != nil {
            return nil, err
        }

        file.Close()
    }

    // create DB and connect
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        logger.Errorf("sqlite opening error: %v", err)
        return nil, err
    }

    // migrate schema
    err = db.AutoMigrate(&schemas.Opening{})
    if err != nil {
        logger.Errorf("sqlite automigrations error: %v", err)
    }

    return db, nil
}
