package database

import (
	"log"
	"shop/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "host=db user=postgres password=postgres dbname=shop port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	DB.AutoMigrate(models.User{})
}
