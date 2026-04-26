package repositories

import (
	"context"
	"fmt"
	"portfolio-backend/internal/core/domain"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type postgresProjectRepository struct {
	db *sqlx.DB
}

// NewPostgresProjectRepository creates a new instance of the project repository for Postgres.
func NewPostgresProjectRepository(db *sqlx.DB) *postgresProjectRepository {
	return &postgresProjectRepository{db: db}
}

func (r *postgresProjectRepository) GetAll(ctx context.Context) ([]domain.Project, error) {
	var projects []domain.Project
	query := `SELECT id, title, description, image_url, github_url, live_url, category, tech_stack, created_at FROM projects ORDER BY created_at DESC`
	
	err := r.db.SelectContext(ctx, &projects, query)
	if err != nil {
		return nil, fmt.Errorf("error getting projects: %w", err)
	}

	return projects, nil
}

func (r *postgresProjectRepository) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	var project domain.Project
	query := `SELECT id, title, description, image_url, github_url, live_url, category, tech_stack, created_at FROM projects WHERE id = $1`

	err := r.db.GetContext(ctx, &project, query, id)
	if err != nil {
		return nil, fmt.Errorf("error getting project by id: %w", err)
	}

	return &project, nil
}

func (r *postgresProjectRepository) Create(ctx context.Context, p *domain.Project) error {
	query := `INSERT INTO projects (title, description, image_url, github_url, live_url, category, tech_stack)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at`

	err := r.db.QueryRowContext(ctx, query, p.Title, p.Description, p.ImageURL, p.GithubURL, p.LiveURL, p.Category, pq.Array(p.TechStack)).
		Scan(&p.ID, &p.CreatedAt)
	
	if err != nil {
		return fmt.Errorf("error creating project: %w", err)
	}

	return nil
}

func (r *postgresProjectRepository) Update(ctx context.Context, p *domain.Project) error {
	query := `UPDATE projects SET title = $1, description = $2, image_url = $3, github_url = $4, live_url = $5, category = $6, tech_stack = $7
	WHERE id = $8`

	_, err := r.db.ExecContext(ctx, query, p.Title, p.Description, p.ImageURL, p.GithubURL, p.LiveURL, p.Category, pq.Array(p.TechStack), p.ID)
	if err != nil {
		return fmt.Errorf("error updating project: %w", err)
	}

	return nil
}

func (r *postgresProjectRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM projects WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting project: %w", err)
	}

	return nil
}
