package controller

import (
	"encoding/json"
	"net/http"

	"github.com/noyandey88/go-todo-app/internal/auth"
	"github.com/noyandey88/go-todo-app/internal/auth/service"
	"github.com/noyandey88/go-todo-app/pkg/response"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService}
}

// SignUp godoc
// @Summary User Sign Up
// @Description Create a new user account
// @Tags auth-controller
// @Accept json
// @Produce json
// @Param data body auth.SignUpRequest true "Sign Up Request"
// @Success 200 {object} user.User
// @Failure 400 {object} response.Response
// @Router /api/auth/sign-up [post]
func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var req auth.SignUpRequest
	json.NewDecoder(r.Body).Decode(&req)
	usr, err := c.authService.SignUp(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.JsonResponse(w, http.StatusOK, true, "Singed up successfully", usr)
}

// SignIn godoc
// @Summary User Sign In
// @Description Authenticate user and return JWT tokens
// @Tags auth-controller
// @Accept json
// @Produce json
// @Param data body auth.SignInRequest true "Sign In Request"
// @Success 200 {object} auth.SignInResponse
// @Failure 401 {object} response.Response
// @Router /api/auth/sign-in [post]
func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	var req auth.SignInRequest

	json.NewDecoder(r.Body).Decode(&req)
	resp, err := c.authService.SignIn(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "signed in successfully", resp)
}

// SignOut godoc
// @Summary User Sign Out
// @Description Invalidate refresh token and log user out
// @Tags auth-controller, internal
// @Accept json
// @Produce json
// @Param data body auth.SignOutRequest true "Sign Out Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/auth/sign-out [post]
func (c *AuthController) SignOut(w http.ResponseWriter, r *http.Request) {
	var req auth.SignOutRequest
	json.NewDecoder(r.Body).Decode(&req)
	if err := c.authService.SignOut(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.JsonResponse(w, http.StatusOK, true, "signed out successfully", nil)
}

// ForgotPassword godoc
// @Summary Forgot Password
// @Description Send reset password email with token
// @Tags auth-controller
// @Accept json
// @Produce json
// @Param data body auth.ForgotPasswordRequest true "Forgot Password Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/auth/forgot-password [post]
func (c *AuthController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req auth.ForgotPasswordRequest
	json.NewDecoder(r.Body).Decode(&req)
	if err := c.authService.ForgotPassword(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.JsonResponse(w, http.StatusOK, true, "An email has been sent to reset your email successfully", nil)
}
