package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/noyandey88/go-todo-app/internal/user"
	"github.com/noyandey88/go-todo-app/internal/user/service"
	"github.com/noyandey88/go-todo-app/middleware"
	"github.com/noyandey88/go-todo-app/pkg/response"
	"github.com/noyandey88/go-todo-app/pkg/utils"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

// GetAllUsers Get All godoc
// @Summary Get All
// @Description Get All Users
// @Tags admin-user-controller
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []user.User
// @Router /admin/users [get]
func (s *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.service.GetAllUsers()

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

	response.JsonResponse(
		w,
		http.StatusOK,
		true,
		"Data loaded successfully",
		users,
	)
}

// GetById godoc
// @Summary Get user by ID
// @Description Retrieves a user by its ID
// @Tags admin-user-controller
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} user.User
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /admin/users/{id} [get]
func (s *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	id, err := strconv.Atoi(idstr)

	if err != nil {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid request id",
			nil,
		)
		return
	}

	usr, err := s.service.GetUserByID(uint(id))
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusNotFound,
			false,
			"User not found",
			nil,
		)
		return
	}

	response.JsonResponse(
		w,
		http.StatusOK,
		true,
		"Users loaded successfully",
		usr,
	)
}

// GetMe godoc
// @Summary Get current user
// @Description Retrieves a user by its ID
// @Tags user-controller
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} user.User
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/me [get]
func (s *UserController) GetMe(w http.ResponseWriter, r *http.Request) {
	usrId, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid request id",
			nil,
		)
		return
	}

	usr, err := s.service.GetUserByID(uint(usrId))
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusNotFound,
			false,
			"User not found",
			nil,
		)
		return
	}

	response.JsonResponse(
		w,
		http.StatusOK,
		true,
		"Users loaded successfully",
		usr,
	)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update user details by ID
// @Tags admin-user-controller
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body user.UpdateRequest true "Updated user data"
// @Security BearerAuth
// @Success 200 {object} user.User
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /admin/user/update/{id} [put]
func (s *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req user.UpdateRequest
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

	// Fetch existing user first (optional but good for checking existence)
	existingUser, err := s.service.GetUserByID(uint(id))
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusNotFound,
			false,
			"User not found",
			nil,
		)
		return
	}

	// Update fields
	existingUser.FirstName = req.FirstName
	existingUser.LastName = req.LastName
	existingUser.UpdatedAt = utils.Epoch()

	if err := s.service.UpdateUser(existingUser); err != nil {
		response.JsonResponse(
			w,
			http.StatusInternalServerError,
			false,
			"Failed to update user due to internal server error",
			nil,
		)
		return
	}

	response.JsonResponse(
		w,
		http.StatusOK,
		true,
		"User updated successfully",
		existingUser,
	)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags admin-user-controller
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /admin/users/delete/{id} [delete]
func (s *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JsonResponse(
			w,
			http.StatusBadRequest,
			false,
			"Invalid user id",
			nil,
		)
		return
	}

	err = s.service.DeleteUser(uint(id))
	if err != nil {
		http.Error(w, "User not found or failed to delete", http.StatusInternalServerError)
		return
	}

	response.JsonResponse(
		w,
		http.StatusOK,
		true,
		"User deleted successfully",
		nil,
	)
}
