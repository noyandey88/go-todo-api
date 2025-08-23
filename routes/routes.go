package routes

import (
	"net/http"

	"github.com/noyandey88/go-todo-app/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {
	manager := middleware.NewManager()

	RegisterSwagger(mux)
	RegisterTodosRoutes(mux, manager)
	RegisterAuthRoutes(mux)
	RegisterUserRoutes(mux)

	mux.Handle("/api/", http.StripPrefix("/api", mux))
}
