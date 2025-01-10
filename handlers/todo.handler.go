// Package handlers This package handles request
package handlers

import (
	"encoding/json"
	"net/http"
	"todo-module/models"
	"todo-module/services"
)

type TodoHandlerInterface interface {
	CreateTodoHandler(w http.ResponseWriter, r *http.Request)
	GetTodos(w http.ResponseWriter, r *http.Request)
	GetSingleTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type TodoHandler struct {
	todoService services.TodoServiceInterface
}

func NewTodoHander(todoService services.TodoServiceInterface) TodoHandlerInterface {
	return &TodoHandler{todoService: todoService}
}

func (s *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Call the createTodoService to write database
	str, err := s.todoService.CreateTodo(r, r.Context())
	if err != nil {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	// Respond with giving success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": str})
}

func (s *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Call the GetTodos service to fetch todos from the database
	todosJSon, err := s.todoService.GetAll(r.Context())
	if err != nil {
		http.Error(w, todosJSon, http.StatusInternalServerError)
		return
	}

	// Respond with the todos as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(todosJSon))
}

func (s *TodoHandler) GetSingleTodo(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload into a Todo struct
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusMethodNotAllowed)
		return
	}

	// Call GetOne to get todo
	todo, err := s.todoService.GetOne(r.Context(), todo.ID)
	if err != nil {
		http.Error(w, "Error: whilet getting todo", http.StatusInternalServerError)
		return
	}

	// Convert todo to json
	response, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "Error while marshalling todo", http.StatusInternalServerError)
		return
	}

	// Respond with the todos as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}

func (s *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload into a Todo struct
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusMethodNotAllowed)
		return
	}

	err := s.todoService.UpdateTodo(r.Context(), todo.ID, todo)
	if err != nil {
		http.Error(w, "Error: while updating todo", http.StatusInternalServerError)
		return
	}

	// Respond with giving success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "OK"})
}

func (s *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload into a Todo struct
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusMethodNotAllowed)
		return
	}

	err := s.todoService.DeleteTodo(r.Context(), todo.ID)
	if err != nil {
		http.Error(w, "Error: while deleting todo", http.StatusInternalServerError)
		return
	}

	// Respond with giving success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "OK"})
}
