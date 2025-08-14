package database

import (
	"github.com/noyandey88/go-todo-app/internal/todo"
	"github.com/noyandey88/go-todo-app/internal/user"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&user.User{}, &todo.Todo{})
}
