package db

import (
	"log"

	"github.com/KadirbekSharau/mywealth-backend/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	dbURL := "postgres://pg:pass@database:5432/crud"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(
		&models.EntityUsers{},
		&models.EntityAccounts{},
		&models.EntityCategories{},
		&models.EntityTransactions{},
	)

	//AccountsDataMigrator(db)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}

func CloseDB(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	database.Close()
}
