package services

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"
)

type skillService struct {
	repo ports.SkillRepository
}

func NewSkillService(repo ports.SkillRepository) ports.SkillService {
	return &skillService{repo: repo}
}

func (s *skillService) GetAll(ctx context.Context) ([]domain.Skill, error) {
	return s.repo.GetAll(ctx)
}

func (s *skillService) GetByID(ctx context.Context, id string) (*domain.Skill, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *skillService) Create(ctx context.Context, skill *domain.Skill) error {
	return s.repo.Create(ctx, skill)
}

func (s *skillService) Update(ctx context.Context, skill *domain.Skill) error {
	return s.repo.Update(ctx, skill)
}

func (s *skillService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
