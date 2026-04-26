package repositories

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type experienceRepository struct {
	db *sqlx.DB
}

func NewExperienceRepository(db *sqlx.DB) ports.ExperienceRepository {
	return &experienceRepository{db: db}
}

func (r *experienceRepository) GetAll(ctx context.Context) ([]domain.Experience, error) {
	var experiences []domain.Experience
	err := r.db.SelectContext(ctx, &experiences, "SELECT * FROM experience ORDER BY start_date DESC")
	return experiences, err
}

func (r *experienceRepository) GetByID(ctx context.Context, id string) (*domain.Experience, error) {
	exp := &domain.Experience{}
	err := r.db.GetContext(ctx, exp, "SELECT * FROM experience WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func (r *experienceRepository) Create(ctx context.Context, exp *domain.Experience) error {
	query := `
		INSERT INTO experience (company, role, description, start_date, end_date, is_current)
		VALUES (:company, :role, :description, :start_date, :end_date, :is_current)`
	_, err := r.db.NamedExecContext(ctx, query, exp)
	return err
}

func (r *experienceRepository) Update(ctx context.Context, exp *domain.Experience) error {
	query := `
		UPDATE experience SET 
			company = :company, 
			role = :role, 
			description = :description, 
			start_date = :start_date, 
			end_date = :end_date, 
			is_current = :is_current
		WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, exp)
	return err
}

func (r *experienceRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM experience WHERE id = $1", id)
	return err
}
