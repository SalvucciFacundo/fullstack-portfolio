package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

// ProjectService defines the application logic for projects.
type ProjectService interface {
	GetAll(ctx context.Context) ([]domain.Project, error)
	GetByID(ctx context.Context, id string) (*domain.Project, error)
	Create(ctx context.Context, project *domain.Project) error
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id string) error
}
