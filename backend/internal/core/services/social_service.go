package services

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"
)

type socialService struct {
	repo ports.SocialRepository
}

func NewSocialService(repo ports.SocialRepository) ports.SocialService {
	return &socialService{repo: repo}
}

func (s *socialService) GetAll(ctx context.Context) ([]domain.SocialLink, error) {
	return s.repo.GetAll(ctx)
}

func (s *socialService) Update(ctx context.Context, social *domain.SocialLink) error {
	return s.repo.Update(ctx, social)
}
