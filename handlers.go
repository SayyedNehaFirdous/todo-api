package main

import (
	"encoding/json"
	"net/http"
)

// Temporary storage (memory)
var todos = []Todo{}
var nextID = 1

// GET /todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// POST /todos
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	// request body se JSON read
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// ID assign
	todo.ID = nextID
	nextID++

	// store in memory
	todos = append(todos, todo)

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	for i, t := range todos {
		if t.ID == id {
			// remove item at index i
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}
