package routes

import (
	"net/http"

	"github.com/noyandey88/go-todo-app/database"
	"github.com/noyandey88/go-todo-app/internal/user/controller"
	"github.com/noyandey88/go-todo-app/internal/user/repository"
	"github.com/noyandey88/go-todo-app/internal/user/service"
	"github.com/noyandey88/go-todo-app/middleware"
	"github.com/noyandey88/go-todo-app/pkg/utils"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	router := utils.NewMuxRouter(mux)
	manager := middleware.NewManager()

	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	todoController := controller.NewUserController(userService)

	router.Get("/users", manager.With(
		middleware.Logger,
		middleware.JWTAuth,
	)(http.HandlerFunc(todoController.GetAllUsers)))

	router.Get("/users/me", manager.With(
		middleware.Logger,
		middleware.JWTAuth,
	)(http.HandlerFunc(todoController.GetMe)))

	router.Get("/users/{id}", manager.With(
		middleware.Logger,
		middleware.JWTAuth,
	)(http.HandlerFunc(todoController.GetById)))

	router.Put("/users/update/{id}", manager.With(
		middleware.Logger,
		middleware.JWTAuth,
	)(http.HandlerFunc(todoController.UpdateUser)))

	router.Delete("/users/delete/{id}", manager.With(
		middleware.Logger,
		middleware.JWTAuth,
	)(http.HandlerFunc(todoController.DeleteUser)))

}
