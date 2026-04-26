package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

type SocialRepository interface {
	GetAll(ctx context.Context) ([]domain.SocialLink, error)
	Update(ctx context.Context, social *domain.SocialLink) error
}
