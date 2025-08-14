package routes

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	RegisterSwagger(mux)
	RegisterTodosRoutes(mux)
	RegisterAuthRoutes(mux)
}
