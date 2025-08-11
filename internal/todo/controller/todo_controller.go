package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/noyandey88/go-todo-app/internal/todo"
	"github.com/noyandey88/go-todo-app/internal/todo/service"
	"github.com/noyandey88/go-todo-app/pkg/response"
)

type TodoController struct {
	service service.TodoService
}

func NewTodoController(service service.TodoService) *TodoController {
	return &TodoController{service: service}
}

// Get All godoc
// @Summary Get All
// @Description Get All Todos
// @Tags todos-controller
// @Accept json
// @Produce json
// @Success 200 {object} []todo.Todo
// @Router /todos [get]
func (c *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	employees, err := c.service.GetAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "Todos loaded successfully", employees)
}

// GetById godoc
// @Summary Get todo by ID
// @Description Retrieves a todo item by its ID
// @Tags todos-controller
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} todo.Todo
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal server error"
// @Router /todos/{id} [get]
func (c *TodoController) GetById(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	id, err := strconv.Atoi(idstr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employees, err := c.service.GetTodoByID(uint(id))
	if err != nil {
		response.JsonResponse(w, http.StatusNotFound, false, "Todo not found", nil)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "Todos loaded successfully", employees)
}

// CreateTodo godoc
// @Summary Create Todos Endpoint
// @Description Creates a new todo item
// @Tags todos-controller
// @Accept json
// @Produce json
// @Param todo body todo.TodoCreateRequest true "create todo"
// @Success 200 {object} todo.Todo
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /todos/create [post]
func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.service.CreateTodo(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "Todo loaded successfully", todo)
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update todo details by ID
// @Tags todos-controller
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body todo.TodoUpdateRequest true "Updated todo data"
// @Success 200 {object} todo.Todo
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal server error"
// @Router /todos/update/{id} [put]
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req todo.TodoUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch existing todo first (optional but good for checking existence)
	existingTodo, err := c.service.GetTodoByID(uint(id))
	if err != nil {
		response.JsonResponse(w, http.StatusNotFound, false, "Todo not found", nil)
		return
	}

	// Update fields
	existingTodo.Title = req.Title
	existingTodo.Description = req.Description
	existingTodo.Completed = req.Completed

	if err := c.service.UpdateTodo(existingTodo); err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "Todo updated successfully", existingTodo)
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo by ID
// @Tags todos-controller
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal server error"
// @Router /todos/delete/{id} [delete]
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.service.DeleteTodo(uint(id))
	if err != nil {
		http.Error(w, "Todo not found or failed to delete", http.StatusInternalServerError)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "Todo deleted successfully", nil)
}
