package services

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"
)

type heroService struct {
	repo ports.HeroRepository
}

func NewHeroService(repo ports.HeroRepository) ports.HeroService {
	return &heroService{repo: repo}
}

func (s *heroService) Get(ctx context.Context) (*domain.HeroSection, error) {
	return s.repo.Get(ctx)
}

func (s *heroService) Update(ctx context.Context, p *domain.HeroSection) error {
	return s.repo.Update(ctx, p)
}
