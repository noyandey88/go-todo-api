package database

import (
	"fmt"
	"log"

	config "github.com/noyandey88/go-todo-app/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
		cfg.Database.TimeZone,
	)

	// dsn := "host=localhost user=postgres password=1234 dbname=go-ecommerce port=5432 sslmode=disable TimeZone=Asia/Dhaka"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB = database
	log.Println("Database connection established.")

	if err := migrate(DB); err != nil {
		panic(err)
	}

	if err := seedSuperAdmin(DB); err != nil {
		panic(err)
	}

	log.Println("Connected to PostgreSQL database using GORM!")
}
