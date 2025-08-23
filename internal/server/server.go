package server

import (
	"fmt"
	"net/http"

	config "github.com/noyandey88/go-todo-app/configs"
	"github.com/noyandey88/go-todo-app/middleware"
	"github.com/noyandey88/go-todo-app/routes"
)

func ConnectServer() {
	mux := http.NewServeMux()
	cfg := config.LoadConfig()
	manager := middleware.NewManager()
	port := fmt.Sprintf(":%d", cfg.Server.Port)

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	// Load all routes
	routes.RegisterRoutes(mux)

	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
