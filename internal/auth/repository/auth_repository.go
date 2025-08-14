package repository

import (
	"github.com/noyandey88/go-todo-app/internal/auth"
	"gorm.io/gorm"
)

type AuthRepository interface {
	SaveBlacklistedToken(token string, expiresAt int64) error
	IsTokenBlacklisted(token string) (bool, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r authRepository) SaveBlacklistedToken(token string, expiresAt int64) error {
	var tokenBlacklist auth.TokenBlacklist
	tokenBlacklist.Token = token
	tokenBlacklist.ExpiresAt = expiresAt
	return r.db.Create(&tokenBlacklist).Error
}

func (r authRepository) IsTokenBlacklisted(token string) (bool, error) {
	var count int64
	err := r.db.Model(&auth.TokenBlacklist{}).Where("token = ?", token).Count(&count).Error
	return count > 0, err
}
