package routes

import (
	"net/http"

	"github.com/noyandey88/go-todo-app/database"
	"github.com/noyandey88/go-todo-app/internal/todo/controller"
	"github.com/noyandey88/go-todo-app/internal/todo/repository"
	"github.com/noyandey88/go-todo-app/internal/todo/service"
)

func RegisterTodosRoutes(mux *http.ServeMux) {
	todoRepo := repository.NewTodoRepository(database.DB)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	mux.Handle("GET /todos", http.HandlerFunc(todoController.GetAllTodos))
	mux.Handle("GET /todos/{id}", http.HandlerFunc(todoController.GetById))
	mux.Handle("POST /todos/create", http.HandlerFunc(todoController.CreateTodo))
	mux.Handle("PUT /todos/update/{id}", http.HandlerFunc(todoController.UpdateTodo))
	mux.Handle("DELETE /todos/delete/{id}", http.HandlerFunc(todoController.DeleteTodo))
}
