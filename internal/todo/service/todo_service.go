package service

import (
	"github.com/noyandey88/go-todo-app/internal/todo"
	"github.com/noyandey88/go-todo-app/internal/todo/repository"
)

type TodoService interface {
	GetAllTodos() ([]todo.Todo, error)
	GetTodoByID(id uint) (*todo.Todo, error)
	CreateTodo(employee *todo.Todo) error
	UpdateTodo(employee *todo.Todo) error
	DeleteTodo(id uint) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetAllTodos() ([]todo.Todo, error) {
	return s.repo.FindAll()
}

func (s *todoService) GetTodoByID(id uint) (*todo.Todo, error) {
	return s.repo.FindByID(id)
}

func (s *todoService) CreateTodo(employee *todo.Todo) error {
	return s.repo.Create(employee)
}

func (s *todoService) UpdateTodo(employee *todo.Todo) error {
	return s.repo.Update(employee)
}

func (s *todoService) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}
