package main

import (
	"github.com/noyandey88/go-todo-app/database"
	"github.com/noyandey88/go-todo-app/internal/server"
)

func main() {
	// connect database
	database.ConnectDatabase()

	// connect server
	server.ConnectServer()
}
