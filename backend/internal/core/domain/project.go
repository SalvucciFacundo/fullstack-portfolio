package domain

import (
	"time"

	"github.com/google/uuid"
)

// Project represents a portfolio project.
type Project struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	ImageURL    string    `db:"image_url" json:"image_url"`
	GithubURL   string    `db:"github_url" json:"github_url"`
	LiveURL     string    `db:"live_url" json:"live_url"`
	Category    string    `db:"category" json:"category"`
	TechStack   []string  `db:"tech_stack" json:"tech_stack"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
