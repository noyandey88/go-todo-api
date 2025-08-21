package database

import (
	"log"

	config "github.com/noyandey88/go-todo-app/configs"
	"github.com/noyandey88/go-todo-app/internal/user"
	"github.com/noyandey88/go-todo-app/pkg/utils"
	"gorm.io/gorm"
)

func seedSuperAdmin(db *gorm.DB) error {
	cfg := config.LoadConfig()

	var existing user.User
	result := db.Where("email = ?", cfg.Database.SuperAdmin).First(&existing)

	if result.Error == nil {
		log.Println("Super admin already exists, skipping seeding.")
		return nil
	}

	password, err := utils.HashPassword(cfg.Database.SuperAdminPassword)

	if err != nil {
		return err
	}

	superAdmin := user.User{
		FirstName: "Super",
		LastName:  "Admin",
		Email:     cfg.Database.SuperAdmin,
		Password:  password,
		Role:      "super_admin",
	}

	if err := db.Create(&superAdmin).Error; err != nil {
		return err
	}

	log.Println("Super admin user created successfully.")

	return nil
}
