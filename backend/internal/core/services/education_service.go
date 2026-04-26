package services

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"
)

type educationService struct {
	repo ports.EducationRepository
}

func NewEducationService(repo ports.EducationRepository) ports.EducationService {
	return &educationService{repo: repo}
}

func (s *educationService) GetAll(ctx context.Context) ([]domain.Education, error) {
	return s.repo.GetAll(ctx)
}

func (s *educationService) GetByID(ctx context.Context, id string) (*domain.Education, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *educationService) Create(ctx context.Context, e *domain.Education) error {
	return s.repo.Create(ctx, e)
}

func (s *educationService) Update(ctx context.Context, e *domain.Education) error {
	return s.repo.Update(ctx, e)
}

func (s *educationService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
