package service

import (
	"github.com/noyandey88/go-todo-app/internal/user"
	"github.com/noyandey88/go-todo-app/internal/user/repository"
)

type UserService interface {
	GetAllUsers() ([]user.User, error)
	GetUserByID(id uint) (*user.User, error)
	UpdateUser(user *user.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (r *userService) GetAllUsers() ([]user.User, error) {
	return r.repo.FindAll()
}

func (r *userService) GetUserByID(id uint) (*user.User, error) {
	return r.repo.FindByID(id)
}

func (r *userService) UpdateUser(user *user.User) error {
	return r.repo.Update(user)
}

func (r *userService) DeleteUser(id uint) error {
	return r.repo.Delete(id)
}
