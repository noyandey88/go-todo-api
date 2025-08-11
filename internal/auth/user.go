package auth

import dbbase "github.com/noyandey88/go-todo-app/internal/db-base"

type User struct {
	dbbase.BaseModel
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required" gorm:"not null"`
}

type SignUpRequest struct {
	FirstName string `json:"firstName" binding:"required" gorm:"not null"`
	LastName  string `json:"lastName" binding:"required" gorm:"not null"`
	Email     string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password  string `json:"password" binding:"required" gorm:"not null"`
	Username  string `json:"username" binding:"required" gorm:"unique;not null"`
}
