package repository

import (
	"github.com/noyandey88/go-todo-app/internal/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]user.User, error)
	FindByID(id uint) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
	Create(user *user.User) error
	Update(user *user.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]user.User, error) {
	var users []user.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepository) FindByID(id uint) (*user.User, error) {
	var user user.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*user.User, error) {
	var user user.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *userRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *user.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&user.User{}, id).Error
}
