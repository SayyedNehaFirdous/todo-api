package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// Routes
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", todoByIDHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Route controller
func todosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getTodos(w, r)
	} else if r.Method == http.MethodPost {
		createTodo(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func todoByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// /todos/1 -> "1"
	idStr := r.URL.Path[len("/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	deleteTodo(w, r, id)
}
