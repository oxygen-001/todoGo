// This package defines logic code to service todo
package services

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-module/models"
	"todo-module/repositories"
)

type TodoServiceInterface interface {
	CreateTodo(r *http.Request, ctx context.Context) (string, error)
	GetAll(ctx context.Context) (string, error)
	GetOne(ctx context.Context, id int) (models.Todo, error)
	UpdateTodo(ctx context.Context, id int, todo models.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}

type TodoService struct {
	repo repositories.TodoRepositoryInterface
}

func NewTodoService(repo repositories.TodoRepositoryInterface) TodoServiceInterface {
	return &TodoService{repo: repo}
}

// WriteTodo Function to write new todo into database

func (s *TodoService) CreateTodo(r *http.Request, ctx context.Context) (string, error) {
	// Decode the JSON payload into a Todo struct
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return "Invalid JSON payload", err
	}

	err := s.repo.CreateTodo(ctx, todo)
	if err != nil {
		return "Error: while inserting to db", err
	}

	return "OK", nil
}

// GetTodos function gets all todos from database

func (s *TodoService) GetAll(ctx context.Context) (string, error) {
	todos, err := s.repo.GetAll(ctx)
	if err != nil {
		return "Error: while getting todos", err
	}

	todosJSON, err := json.Marshal(todos)
	if err != nil {
		return "Error: while parsing to json", err
	}

	return string(todosJSON), nil
}

func (s *TodoService) GetOne(ctx context.Context, id int) (models.Todo, error) {
	return s.repo.GetOne(ctx, id)
}

func (s *TodoService) UpdateTodo(ctx context.Context, id int, todo models.Todo) error {
	return s.repo.UpdateTodo(ctx, id, todo)
}

func (s *TodoService) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.DeleteTodo(ctx, id)
}
