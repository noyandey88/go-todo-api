package routes

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	RegisterSwagger(mux)
	RegisterTodosRoutes(mux)
	RegisterAuthRoutes(mux)
	RegisterUserRoutes(mux)

	mux.Handle("/api/", http.StripPrefix("/api", mux))
}
