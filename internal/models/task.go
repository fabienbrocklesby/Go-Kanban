package models

import "time"

type Task struct {
		ID               int
		Title            string
		Description      string
		Status           string
		GitHubRepo       string
		GitHubIssueNumber int
		CreatedAt        time.Time
}