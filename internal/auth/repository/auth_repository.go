package repository

import (
	"gorm.io/gorm"
)

type AuthRepository interface {
	FSaveBlacklistedToken(token string, expiresAt int64) error
	IsTokenBlacklisted(token string) (bool, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r authRepository) SaveBlacklistedToken(token string, expiresAt int64) error {
	return 
}