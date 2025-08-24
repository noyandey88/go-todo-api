package main

import (
	"github.com/noyandey88/go-todo-app/database"
	_ "github.com/noyandey88/go-todo-app/docs"
	"github.com/noyandey88/go-todo-app/internal/server"
)

// @title Go Todos API
// @version 1.0
// @description This is a sample server Todos server.
// @termsOfService http://swagger.io/terms/new

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host go-todo-api-36ta.onrender.com
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide the access token directly, without "Bearer ".
func main() {
	// connect database
	database.ConnectDatabase()

	// connect server
	server.ConnectServer()
}
