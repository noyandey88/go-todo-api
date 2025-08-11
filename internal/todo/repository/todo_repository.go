package repository

import (
	"github.com/noyandey88/go-todo-app/internal/todo"
	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll() ([]todo.Todo, error)
	FindByID(id uint) (*todo.Todo, error)
	Create(employee *todo.Todo) error
	Update(employee *todo.Todo) error
	Delete(id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) FindAll() ([]todo.Todo, error) {
	var todos []todo.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *todoRepository) FindByID(id uint) (*todo.Todo, error) {
	var todo todo.Todo
	result := r.db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (r *todoRepository) Create(todo *todo.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) Update(todo *todo.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&todo.Todo{}, id).Error
}
