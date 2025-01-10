package main

// TODO: Check migration if table doesn't exit then enter db.go and uncomment RunMigration function

import (
	"fmt"
	"log"
	"net/http"
	"todo-module/handlers"
	"todo-module/pkg"
	"todo-module/repositories"
	"todo-module/services"
)

func main() {
	// Initiate database
	db, err := pkg.InitDB()
	if err != nil {
		// Log a fatal error and stop execution if the database connection fails
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Initialize the repository
	todoRepo := &repositories.TodoRepository{DB: db}

	// Initialize the service
	todoService := services.NewTodoService(todoRepo)

	// Initialize the handler
	todoHandler := handlers.NewTodoHander(todoService)

	// Close database connection when the program exits
	defer db.Close()

	http.HandleFunc("GET /", todoHandler.GetTodos)
	http.HandleFunc("GET /single", todoHandler.GetSingleTodo)
	http.HandleFunc("POST /create", todoHandler.CreateTodoHandler)
	http.HandleFunc("PUT /", todoHandler.UpdateTodo)
	http.HandleFunc("DELETE /", todoHandler.DeleteTodo)

	fmt.Println("Server is running on http://localhost:8080")

	// Use `ListenAndServe` to start the server on port 8080 and handle requests
	log.Fatal(http.ListenAndServe(":8080", nil))
}
