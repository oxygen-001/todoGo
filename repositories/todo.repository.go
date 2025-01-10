// Package repositories defines all database actions
package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"todo-module/models"
)

// region Declaration

type TodoRepositoryInterface interface {
	// Get
	GetAll(ctx context.Context) ([]models.Todo, error)
	GetOne(ctx context.Context, id int) (models.Todo, error)
	// Create
	CreateTodo(ctx context.Context, todo models.Todo) error
	// Update
	UpdateTodo(ctx context.Context, id int, todo models.Todo) error
	// Delete
	DeleteTodo(ctx context.Context, id int) error
}

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(DB *sql.DB) TodoRepositoryInterface {
	return &TodoRepository{DB: DB}
}

// endregion

// region Get

func (r *TodoRepository) GetAll(ctx context.Context) ([]models.Todo, error) {
	// Create a slice
	var todos = make([]models.Todo, 0)

	// Prepare query
	query := `SELECT id, title, done, created_at FROM todos`

	// Get todos using context
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return todos, err
	}

	defer rows.Close()

	fmt.Println(rows)

	// Iterate over rows and append them into todos
	for rows.Next() {
		var todo models.Todo

		// Scan row into todo struct
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.Created_at); err != nil {
			return todos, fmt.Errorf("Failed to scan row: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

func (r *TodoRepository) GetOne(ctx context.Context, id int) (models.Todo, error) {
	// Preapare query
	query := `SELECT id, title, done, created_at FROM todos WHERE id = $1`

	var todo models.Todo

	// Use QueryRowContext to get only single todo
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&todo.ID, &todo.Title, &todo.Done, &todo.Created_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Todo{}, fmt.Errorf("todo with id %d not found", id)
		}

		// Handle other types of errors
		return models.Todo{}, err
	}

	return todo, nil
}

//endregion

// Region Create
func (r *TodoRepository) CreateTodo(ctx context.Context, todo models.Todo) error {
	// Prepare query
	query := `INSERT INTO todos (title, done) VALUES ($1, $2)`

	// Execute query
	_, err := r.DB.ExecContext(ctx, query, todo.Title, todo.Done)
	if err != nil {
		return err
	}

	return nil
}

//endregion

// Region Update

func (r *TodoRepository) UpdateTodo(ctx context.Context, id int, todo models.Todo) error {
	// Prepare query
	query := `UPDATE todos SET title = $1, done = $2 WHERE id = $3`

	// Execute query
	result, err := r.DB.ExecContext(ctx, query, todo.Title, todo.Done, id)
	if err != nil {
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Make sure the todo with id was found
	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", id)
	}

	return nil
}

// endregion

// Region Delete

func (r *TodoRepository) DeleteTodo(ctx context.Context, id int) error {
	// Prepare query
	query := `DELETE FROM todos WHERE id = $1`

	// Execute qeury
	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", id)
	}

	return nil
}

// endregion
