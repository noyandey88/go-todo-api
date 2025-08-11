package todo

import dbbase "github.com/noyandey88/go-todo-app/internal/db-base"

type Todo struct {
	dbbase.BaseModel
	Title       string `json:"title" gorm:"not null;unique"`
	Description string `json:"description" gorm:"not null"`
	Completed   bool   `json:"completed" gorm:"default:false"`
	// CreatedBy      string `json:"createdBy" gorm:"not null"`
	// LastModifiedBy string `json:"lastModifiedBy" gorm:"not null"`
}

type TodoCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type TodoUpdateRequest struct {
	ID          uint   `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed" validate:"required"`
}
