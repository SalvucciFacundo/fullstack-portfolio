package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

// ExperienceService defines the application logic for experiences.
type ExperienceService interface {
	GetAll(ctx context.Context) ([]domain.Experience, error)
	GetByID(ctx context.Context, id string) (*domain.Experience, error)
	Create(ctx context.Context, e *domain.Experience) error
	Update(ctx context.Context, e *domain.Experience) error
	Delete(ctx context.Context, id string) error
}

// EducationService defines the application logic for education entries.
type EducationService interface {
	GetAll(ctx context.Context) ([]domain.Education, error)
	GetByID(ctx context.Context, id string) (*domain.Education, error)
	Create(ctx context.Context, e *domain.Education) error
	Update(ctx context.Context, e *domain.Education) error
	Delete(ctx context.Context, id string) error
}

// HeroService defines the application logic for the hero section.
type HeroService interface {
	Get(ctx context.Context) (*domain.HeroSection, error)
	Update(ctx context.Context, p *domain.HeroSection) error
}

// SkillService defines the application logic for skills.
type SkillService interface {
	GetAll(ctx context.Context) ([]domain.Skill, error)
	GetByID(ctx context.Context, id string) (*domain.Skill, error)
	Create(ctx context.Context, s *domain.Skill) error
	Update(ctx context.Context, s *domain.Skill) error
	Delete(ctx context.Context, id string) error
}

// SocialService defines the application logic for social links.
type SocialService interface {
	GetAll(ctx context.Context) ([]domain.SocialLink, error)
	Update(ctx context.Context, s *domain.SocialLink) error
}
