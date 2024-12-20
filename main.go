package main

// TODO: Check migration if table doesn't exit then enter db.go and uncomment RunMigration function

import (
	"fmt"
	"log"
	"net/http"
	"todo-module/handlers"
	"todo-module/pkg"
)

func main() {
	// Initiate database
	db, err := pkg.InitDB()
	if err != nil {
		// Log a fatal error and stop execution if the database connection fails
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Close database connection when the program exits
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to-do app")
	})

	http.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTodoHandler(r, w, db)
	})

	http.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTodosHandler(r, w, db)
	})

	fmt.Println("Server is running on http://localhost:8080")

	// Use `ListenAndServe` to start the server on port 8080 and handle requests
	log.Fatal(http.ListenAndServe(":8080", nil))
}
