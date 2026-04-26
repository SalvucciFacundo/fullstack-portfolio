package services

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"
)

type experienceService struct {
	repo ports.ExperienceRepository
}

func NewExperienceService(repo ports.ExperienceRepository) ports.ExperienceService {
	return &experienceService{repo: repo}
}

func (s *experienceService) GetAll(ctx context.Context) ([]domain.Experience, error) {
	return s.repo.GetAll(ctx)
}

func (s *experienceService) GetByID(ctx context.Context, id string) (*domain.Experience, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *experienceService) Create(ctx context.Context, e *domain.Experience) error {
	return s.repo.Create(ctx, e)
}

func (s *experienceService) Update(ctx context.Context, e *domain.Experience) error {
	return s.repo.Update(ctx, e)
}

func (s *experienceService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
