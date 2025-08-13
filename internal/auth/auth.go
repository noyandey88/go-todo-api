package auth

import "github.com/noyandey88/go-todo-app/internal/user"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=6"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=6"`
}

type SignUpRequest struct {
	FirstName string `json:"firstName" binding:"required" gorm:"not null"`
	LastName  string `json:"lastName" binding:"required" gorm:"not null"`
	Email     string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password  string `json:"password" binding:"required,min=6" gorm:"not null"`
}

type SignOutRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type SignInResponse struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	TokenType    string    `json:"tokenType"`
	User         user.User `json:"user"`
}
