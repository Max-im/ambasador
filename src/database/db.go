package database

import (
	"log"
	"shop/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "host=db user=postgres password=postgres dbname=shop port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
		return err
	}
	// Auto migrate your model structs
	AutoMigrate()
	return nil
}

func AutoMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Link{},
		&models.Order{},
		&models.OrderItem{},
	)
}
