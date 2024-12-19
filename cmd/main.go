package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fabienbrocklesby/Go-Kanban/internal/database"
	"github.com/fabienbrocklesby/Go-Kanban/internal/models"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello, world!")
}

func createTask(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := database.CreateTask(
        task.Title,
        task.Description,
        task.Status,
        task.GitHubRepo,
        task.GitHubIssueNumber,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully"})
}

func main() {
    err := database.InitDB()
    if err != nil {
        fmt.Printf("Error connecting to database: %s\n", err)
        return
    }
    defer database.DB.Close()

    http.HandleFunc("/", getRoot)
    http.HandleFunc("/tasks", createTask)

    fmt.Println("Server starting on http://localhost:4000")
    err = http.ListenAndServe(":4000", nil)
    if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}