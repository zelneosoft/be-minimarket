package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Output log ke terminal
		logger.Config{
			SlowThreshold: time.Second, // Menampilkan log untuk query lambat
			LogLevel:      logger.Info, // Menampilkan semua query
			Colorful:      true,        // Menampilkan log dengan warna
		},
	)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // Tambahkan logger ke konfigurasi GORM
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Database connection successfully opened")
}
