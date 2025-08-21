package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/noyandey88/go-todo-app/internal/todo"
	"github.com/noyandey88/go-todo-app/internal/todo/service"
	"github.com/noyandey88/go-todo-app/middleware"
	"github.com/noyandey88/go-todo-app/pkg/response"
)

type TodoController struct {
	service service.TodoService
}

func NewTodoController(service service.TodoService) *TodoController {
	return &TodoController{service: service}
}

// GetAllTodos Get All godoc
// @Summary Get All
// @Description Get All Todos
// @Tags todo-controller
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []todo.Todo
// @Router /todos [get]
func (c *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := c.service.GetAllTodos()

	if err != nil {
		response.JsonResponse(
			w,
			http.StatusInternalServerError,
			false,
			"Internal Server Error",
			nil,
		)
		return
	}

	usrId, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid accessToken or user id",
			nil,
		)
		return
	}

	var filteredTodo []todo.Todo

	for _, t := range todos {
		if t.UserId == usrId {
			filteredTodo = append(filteredTodo, t)
		}
	}

	response.JsonResponse(
		w,
		http.StatusOK,
		true,
		"Todos loaded successfully",
		filteredTodo,
	)
}

// GetById godoc
// @Summary Get todo by ID
// @Description Retrieves a todo item by its ID
// @Tags todo-controller
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Security BearerAuth
// @Success 200 {object} todo.Todo
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal server error"
// @Router /todos/{id} [get]
func (c *TodoController) GetById(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	id, err := strconv.Atoi(idstr)

	if err != nil {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Internal request id",
			nil,
		)
		return
	}

	usrId, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid accessToken or user id",
			nil,
		)
		return
	}

	todo, err := c.service.GetTodoByID(uint(id))
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusNotFound,
			false,
			"Todo not found",
			nil,
		)
		return
	}

	if todo.UserId != usrId {
		response.JsonResponse(w,
			http.StatusNotFound,
			false,
			"Todo not found",
			nil,
		)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "Todos loaded successfully", todo)
}

// CreateTodo godoc
// @Summary Create Todos Endpoint
// @Description Creates a new todo item
// @Tags todo-controller
// @Accept json
// @Produce json
// @Param todo body todo.TodoCreateRequest true "create todo"
// @Security BearerAuth
// @Success 200 {object} todo.Todo
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /todos/create [post]
func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo todo.Todo

	usrId, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid accessToken or user id",
			nil,
		)
		return
	}

	todo.UserId = usrId

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		response.JsonResponse(
			w,
			http.StatusInternalServerError,
			false,
			"Failed to decode to json",
			nil,
		)
		return
	}

	if err := c.service.CreateTodo(&todo); err != nil {
		response.JsonResponse(
			w,
			http.StatusInternalServerError,
			false,
			"Failed to create todo",
			nil,
		)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "Todo loaded successfully", todo)
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update todo details by ID
// @Tags todo-controller
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body todo.TodoUpdateRequest true "Updated todo data"
// @Security BearerAuth
// @Success 200 {object} todo.Todo
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal server error"
// @Router /todos/update/{id} [put]
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid ID",
			nil,
		)
		return
	}

	usrId, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid accessToken or user id",
			nil,
		)
		return
	}

	var req todo.TodoUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid request body",
			nil,
		)
		return
	}

	// Fetch existing todo first (optional but good for checking existence)
	existingTodo, err := c.service.GetTodoByID(uint(id))
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusNotFound,
			false,
			"Todo not found",
			nil,
		)
		return
	}

	if existingTodo.UserId != usrId {
		response.JsonResponse(w, http.StatusNotFound, false, "Todo not found", nil)
		return
	}

	// Update fields
	existingTodo.Title = req.Title
	existingTodo.Description = req.Description
	existingTodo.Completed = req.Completed

	if err := c.service.UpdateTodo(existingTodo); err != nil {
		response.JsonResponse(
			w,
			http.StatusInternalServerError,
			false,
			"Failed to update json",
			nil,
		)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "Todo updated successfully", existingTodo)
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo by ID
// @Tags todo-controller
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Security BearerAuth
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Todo not found"
// @Failure 500 {string} string "Internal server error"
// @Router /todos/delete/{id} [delete]
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"invalid request id",
			nil,
		)
		return
	}

	usrId, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid accessToken or user id",
			nil,
		)
		return
	}

	// Fetch existing todo first (optional but good for checking existence)
	existingTodo, err := c.service.GetTodoByID(uint(id))
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusNotFound,
			false,
			"Todo not found",
			nil,
		)
		return
	}

	if existingTodo.UserId != usrId {
		response.JsonResponse(
			w,
			http.StatusNotFound,
			false,
			"Todo not found",
			nil,
		)
		return
	}

	err = c.service.DeleteTodo(uint(id))
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusInternalServerError,
			false,
			"Todo not found or failed to delete",
			nil,
		)
		return
	}

	response.JsonResponse(
		w,
		http.StatusOK,
		true,
		"Todo deleted successfully",
		nil,
	)
}
