package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

// ProjectRepository defines the contract for project persistence.
type ProjectRepository interface {
	GetAll(ctx context.Context) ([]domain.Project, error)
	GetByID(ctx context.Context, id string) (*domain.Project, error)
	Create(ctx context.Context, project *domain.Project) error
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id string) error
}
