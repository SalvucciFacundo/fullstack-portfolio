package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

type EducationRepository interface {
	GetAll(ctx context.Context) ([]domain.Education, error)
	GetByID(ctx context.Context, id string) (*domain.Education, error)
	Create(ctx context.Context, education *domain.Education) error
	Update(ctx context.Context, education *domain.Education) error
	Delete(ctx context.Context, id string) error
}
