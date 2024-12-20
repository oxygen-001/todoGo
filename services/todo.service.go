// This package defines logic code to service todo
package services

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todo-module/models"
)

// WriteTodo Function to write new todo into database
func WriteTodo(r *http.Request, db *sql.DB) (string, error) {

	// Decode the JSON payload into a Todo struct
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return "Invalid JSON payload", err
	}

	// Query to insert new Todo
	query := `INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING id`

	// Execute query
	err := db.QueryRow(query, todo.Title, todo.Done).Scan(&todo.ID)
	if err != nil {
		return "Error while writing todo into db", err
	}

	return "OK", nil
}

// GetTodos function gets all todos from database
func GetTodos(r *http.Request, db *sql.DB) (string, error) {

	// Query to get all todos
	query := `SELECT id, title, done, created_at FROM todos`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return "Error while fetching todos from database", err
	}
	defer rows.Close() // Ensure rows are properly closed to avoid resource leaks

	// Slice to hold todos
	var todos []models.Todo

	// Iterate through rows
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.Created_at); err != nil {
			return "Error while scanning todo rows", err
		}

		// Push todo to todos
		todos = append(todos, todo)
	}

	// Check if an error occurred during iteration
	if err := rows.Err(); err != nil {
		return "Error while iterating over rows", err
	}

	// Convert todos slice to JSON
	todosJSON, err := json.Marshal(todos)
	if err != nil {
		return "Error while encoding todos to JSON", err
	}

	// string(todosJSON) converts the []byte to string for returning as function output
	return string(todosJSON), nil
}
