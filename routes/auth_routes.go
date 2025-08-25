package routes

import (
	"net/http"

	"github.com/noyandey88/go-todo-app/database"
	"github.com/noyandey88/go-todo-app/internal/auth/controller"
	"github.com/noyandey88/go-todo-app/internal/auth/repository"
	"github.com/noyandey88/go-todo-app/internal/auth/service"
	userRepository "github.com/noyandey88/go-todo-app/internal/user/repository"
	"github.com/noyandey88/go-todo-app/middleware"
	"github.com/noyandey88/go-todo-app/pkg/utils"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	router := utils.NewMuxRouter(mux)
	manager := middleware.NewManager()

	authRepo := repository.NewAuthRepository(database.DB)
	userRepo := userRepository.NewUserRepository(database.DB)
	authService := service.NewAuthService(authRepo, userRepo)
	authController := controller.NewAuthController(authService)

	router.Post("/auth/sign-in", manager.With(
		http.HandlerFunc(authController.SignIn),
	))

	router.Post("/auth/sign-up", manager.With(
		http.HandlerFunc(authController.SignUp),
	))

	// router.Post("/auth/sign-out", manager.With(
	// 	http.HandlerFunc(authController.SignOut),
	// ))

	router.Post("/auth/forgot-password", manager.With(
		http.HandlerFunc(authController.ForgotPassword),
	))
}
