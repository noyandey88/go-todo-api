package server

import (
	"fmt"
	"net/http"

	config "github.com/noyandey88/go-todo-app/configs"
	"github.com/noyandey88/go-todo-app/middleware"
)

func ConnectServer() {
	mux := http.NewServeMux()
	cfg := config.LoadConfig()
	port := fmt.Sprintf(":%d", cfg.Server.Port)

	globalRouter := middleware.GlobalRouter(mux)

	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(port, globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
