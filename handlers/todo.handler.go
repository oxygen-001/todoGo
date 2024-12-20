// Package handlers This package handles request
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todo-module/services"
)

func CreateTodoHandler(r *http.Request, w http.ResponseWriter, db *sql.DB) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Call the createTodoService to write database
	str, err := services.WriteTodo(r, db)
	if err != nil {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	// Respond with giving success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": str})
}

func GetTodosHandler(r *http.Request, w http.ResponseWriter, db *sql.DB) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Call the GetTodos service to fetch todos from the database
	todosJSon, err := services.GetTodos(r, db)
	if err != nil {
		http.Error(w, todosJSon, http.StatusInternalServerError)
		return
	}

	// Respond with the todos as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(todosJSon))
}
