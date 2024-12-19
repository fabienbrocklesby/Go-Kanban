package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
    var err error
    DB, err = sql.Open("sqlite3", "./tasks.db")
    if err != nil {
        return fmt.Errorf("failed to open database: %w", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %w", err)
    }

    _, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            description TEXT,
            status TEXT NOT NULL DEFAULT 'todo',
            github_repo TEXT,
            github_issue_number INTEGER,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    return err
}

func CreateTask(title, description, status, githubRepo string, githubIssueNumber int) error {
    query := `
        INSERT INTO tasks (title, description, status, github_repo, github_issue_number)
        VALUES (?, ?, ?, ?, ?)
    `
    _, err := DB.Exec(query, title, description, status, githubRepo, githubIssueNumber)
    return err
}