package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

type HeroRepository interface {
	Get(ctx context.Context) (*domain.HeroSection, error)
	Update(ctx context.Context, hero *domain.HeroSection) error
}
