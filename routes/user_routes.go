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

	manager.Use(
		middleware.Logger,
		middleware.JWTAuth,
	)

	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	todoController := controller.NewUserController(userService)

	router.Get("/admin/users", manager.With(
		http.HandlerFunc(todoController.GetAllUsers),
	))

	router.Get("/users/me", manager.With(
		http.HandlerFunc(todoController.GetMe),
	))

	router.Get("/admin/users/{id}", manager.With(
		http.HandlerFunc(todoController.GetById),
	))

	router.Put("/admin/users/update/{id}", manager.With(
		http.HandlerFunc(todoController.UpdateUser),
	))

	router.Delete("/admin/users/delete/{id}", manager.With(
		http.HandlerFunc(todoController.DeleteUser),
	))
}
