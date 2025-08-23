package routes

import (
	"net/http"

	"github.com/noyandey88/go-todo-app/database"
	"github.com/noyandey88/go-todo-app/internal/todo/controller"
	"github.com/noyandey88/go-todo-app/internal/todo/repository"
	"github.com/noyandey88/go-todo-app/internal/todo/service"
	"github.com/noyandey88/go-todo-app/middleware"
	"github.com/noyandey88/go-todo-app/pkg/utils"
)

func RegisterTodosRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	router := utils.NewMuxRouter(mux)

	manager.Use(middleware.JWTAuth)

	todoRepo := repository.NewTodoRepository(database.DB)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	router.Get("/todos", manager.With(
		http.HandlerFunc(todoController.GetAllTodos),
	))

	router.Get("/todos/{id}", manager.With(
		http.HandlerFunc(todoController.GetById),
	))

	router.Post("/todos/create", manager.With(
		http.HandlerFunc(todoController.CreateTodo),
	))

	router.Put("/todos/update/{id}", manager.With(
		http.HandlerFunc(todoController.UpdateTodo),
	))

	router.Delete("/todos/delete/{id}", manager.With(
		http.HandlerFunc(todoController.DeleteTodo),
	))
}
