package service

import (
	"errors"
	"time"

	config "github.com/noyandey88/go-todo-app/configs"
	"github.com/noyandey88/go-todo-app/internal/auth"
	authRepository "github.com/noyandey88/go-todo-app/internal/auth/repository"
	"github.com/noyandey88/go-todo-app/internal/user"
	userRepository "github.com/noyandey88/go-todo-app/internal/user/repository"
	"github.com/noyandey88/go-todo-app/pkg/jwtutil"
	"github.com/noyandey88/go-todo-app/pkg/utils"
)

type AuthService interface {
	SignUp(req auth.SignUpRequest) (*user.User, error)
	SignIn(req auth.SignInRequest) (*auth.SignInResponse, error)
	SignOut(req auth.SignOutRequest) error
	ForgotPassword(req auth.ForgotPasswordRequest) error
	// ResetPassword(req auth.ResetPasswordRequest) error
	// ChangePassword(userID uint, req auth.ChangePasswordRequest) error
}

type authService struct {
	authRepo authRepository.AuthRepository
	userRepo userRepository.UserRepository
}

func NewAuthService(authRepo authRepository.AuthRepository, userRepo userRepository.UserRepository) AuthService {
	return &authService{authRepo, userRepo}
}

func (s *authService) SignUp(req auth.SignUpRequest) (*user.User, error) {
	var user user.User
	hashedPass, _ := utils.HashPassword(req.Password)

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Password = hashedPass

	err := s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *authService) SignIn(req auth.SignInRequest) (*auth.SignInResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("User not found")
	}

	if utils.CompareHashedPassword(user.Password, req.Password) != nil {
		return nil, errors.New("Invalid credentials")
	}

	accessToken, _ := jwtutil.GenerateAccessToken(user.ID, config.AppConfig.JWT.Secret, config.AppConfig.JWT.ExpiresIn)
	refreshToken, _ := jwtutil.GenerateRefreshToken(user.ID, config.AppConfig.JWT.Secret)

	return &auth.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		User:         *user,
	}, nil
}

func (s *authService) SignOut(req auth.SignOutRequest) error {
	exp := time.Now().Add(7 * 24 * time.Hour).Unix() // match refresh token life
	return s.authRepo.SaveBlacklistedToken(req.RefreshToken, exp)
}

func (s *authService) ForgotPassword(req auth.ForgotPasswordRequest) error {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return errors.New("user not found")
	}
	resetToken, _ := jwtutil.GenerateResetToken(user.ID, config.AppConfig.JWT.Secret)
	// send resetToken via email here
	_ = resetToken
	return nil
}

// func (s *authService) ResetPassword(req auth.ResetPasswordRequest) error {
// 	userID, err := jwtutil.ParseResetToken(req.Token, config.AppConfig.JWT.Secret)
// 	if err != nil {
// 		return errors.New("invalid or expired token")
// 	}
// 	hashedPass, _ := utils.HashPassword(req.NewPassword)
// 	return s.userRepo.UpdatePassword(userID, string(hashedPass))
// }

// func (s *authService) ChangePassword(userID uint, req model.ChangePasswordRequest) error {
// 	user, _ := s.userRepo.GetByID(userID)
// 	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)) != nil {
// 		return errors.New("old password does not match")
// 	}
// 	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
// 	return s.userRepo.UpdatePassword(userID, string(hashedPass))
// }
