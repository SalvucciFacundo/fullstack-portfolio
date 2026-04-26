package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

type SkillRepository interface {
	GetAll(ctx context.Context) ([]domain.Skill, error)
	GetByID(ctx context.Context, id string) (*domain.Skill, error)
	Create(ctx context.Context, skill *domain.Skill) error
	Update(ctx context.Context, skill *domain.Skill) error
	Delete(ctx context.Context, id string) error
}
