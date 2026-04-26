package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

// ExperienceRepository defines the contract for experience persistence.
type ExperienceRepository interface {
	GetAll(ctx context.Context) ([]domain.Experience, error)
	GetByID(ctx context.Context, id string) (*domain.Experience, error)
	Create(ctx context.Context, experience *domain.Experience) error
	Update(ctx context.Context, experience *domain.Experience) error
	Delete(ctx context.Context, id string) error
}
