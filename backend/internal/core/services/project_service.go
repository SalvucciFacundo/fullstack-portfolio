package services

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"
)

type projectService struct {
	repo ports.ProjectRepository
}

// NewProjectService creates a new instance of the project application service.
func NewProjectService(repo ports.ProjectRepository) *projectService {
	return &projectService{repo: repo}
}

func (s *projectService) GetAll(ctx context.Context) ([]domain.Project, error) {
	return s.repo.GetAll(ctx)
}

func (s *projectService) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *projectService) Create(ctx context.Context, p *domain.Project) error {
	// Aquí podrías agregar lógica de negocio adicional si fuera necesaria
	return s.repo.Create(ctx, p)
}

func (s *projectService) Update(ctx context.Context, p *domain.Project) error {
	return s.repo.Update(ctx, p)
}

func (s *projectService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
