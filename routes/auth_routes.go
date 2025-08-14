package routes

import (
	"net/http"

	"github.com/noyandey88/go-todo-app/database"
	"github.com/noyandey88/go-todo-app/internal/auth/controller"
	"github.com/noyandey88/go-todo-app/internal/auth/repository"
	"github.com/noyandey88/go-todo-app/internal/auth/service"
	userRepository "github.com/noyandey88/go-todo-app/internal/user/repository"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	authRepo := repository.NewAuthRepository(database.DB)
	userRepo := userRepository.NewUserRepository(database.DB)
	authService := service.NewAuthService(authRepo, userRepo)
	authController := controller.NewAuthController(authService)

	mux.Handle("POST /auth/signin", http.HandlerFunc(authController.SignIn))
	mux.Handle("POST /auth/signup", http.HandlerFunc(authController.SignUp))
	mux.Handle("POST /auth/signout", http.HandlerFunc(authController.SignOut))
	mux.Handle("POST /auth/forgot-password", http.HandlerFunc(authController.ForgotPassword))
}
