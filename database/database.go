package database

import (
	"log"

	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Rt0011rw007@tcp(localhost:3306)/struct_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Menjalankan migrasi
	if err := database.AutoMigrate(&entity.Product{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
}
