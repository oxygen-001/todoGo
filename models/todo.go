// Package models defines the structure of Todo model
package models

import "time"

// Todo represents a task in the todo list
type Todo struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Done       bool      `json:"done"`
	Created_at time.Time `json:"created_at"`
}
