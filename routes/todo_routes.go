package routes

import (
	"net/http"

	"github.com/noyandey88/go-todo-app/database"
	"github.com/noyandey88/go-todo-app/internal/todo/controller"
	"github.com/noyandey88/go-todo-app/internal/todo/repository"
	"github.com/noyandey88/go-todo-app/internal/todo/service"
	"github.com/noyandey88/go-todo-app/middleware"
)

func RegisterTodosRoutes(mux *http.ServeMux) {
	todoRepo := repository.NewTodoRepository(database.DB)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	mux.Handle("GET /todos", middleware.JWTAuth(http.HandlerFunc(todoController.GetAllTodos)))
	mux.Handle("GET /todos/{id}", middleware.JWTAuth(http.HandlerFunc(todoController.GetById)))
	mux.Handle("POST /todos/create", middleware.JWTAuth(http.HandlerFunc(todoController.CreateTodo)))
	mux.Handle("PUT /todos/update/{id}", middleware.JWTAuth(http.HandlerFunc(todoController.UpdateTodo)))
	mux.Handle("DELETE /todos/delete/{id}", middleware.JWTAuth(http.HandlerFunc(todoController.DeleteTodo)))
}
