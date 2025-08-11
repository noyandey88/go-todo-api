package main

import (
	"github.com/noyandey88/go-todo-app/database"
	_ "github.com/noyandey88/go-todo-app/docs"
	"github.com/noyandey88/go-todo-app/internal/server"
)

// @title Go Todo API
// @Router PATH [get]
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /
func main() {
	// connect database
	database.ConnectDatabase()

	// connect server
	server.ConnectServer()
}
