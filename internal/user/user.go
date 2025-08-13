package user

import dbbase "github.com/noyandey88/go-todo-app/internal/db-base"

type User struct {
	dbbase.BaseModel
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"`
}
