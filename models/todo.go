// Package models defines the structure of Todo model
package models

// Todo represents a task in the todo list
type Todo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Done       bool   `json:"done"`
	Created_at string `json:"created_at"`
}
