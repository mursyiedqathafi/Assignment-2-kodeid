package database

import (
	"assignment-2/config"
	"assignment-2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() (Database, error) {
	dbInfo := config.GetDatabaseEnv()

	var config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Name)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("Error open connection to db", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		fmt.Println("Error on migration")
		return Database{}, err
	}

	return Database{
		db: db,
	}, nil
}
