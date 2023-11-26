package db

import (
	"LK_back/pkg/models"
	"LK_back/settings"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DB() *gorm.DB {
	DBConfig := settings.InitConfigDB()
	dsn := fmt.Sprintf("host=database_lk user=%s password=%s dbname=%s port=5432 sslmode=disable", DBConfig.DBUser, DBConfig.DBPassword, DBConfig.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)

	}
	db.AutoMigrate(&models.User{})
	return db
}
