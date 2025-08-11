package database

import (
	"github.com/noyandey88/go-todo-app/internal/auth"
	"github.com/noyandey88/go-todo-app/internal/todo"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&auth.User{}, &todo.Todo{})
}
