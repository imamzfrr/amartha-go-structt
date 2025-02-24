package database

import (
	"log"

	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user:password@tcp(localhost:3306)/yourdb?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&entity.CategoryEntity{}) // Migrasi tabel category
	DB = database
}
