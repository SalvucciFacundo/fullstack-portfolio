package repositories

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type educationRepository struct {
	db *sqlx.DB
}

func NewEducationRepository(db *sqlx.DB) ports.EducationRepository {
	return &educationRepository{db: db}
}

func (r *educationRepository) GetAll(ctx context.Context) ([]domain.Education, error) {
	var education []domain.Education
	err := r.db.SelectContext(ctx, &education, "SELECT * FROM education ORDER BY start_date DESC")
	return education, err
}

func (r *educationRepository) GetByID(ctx context.Context, id string) (*domain.Education, error) {
	edu := &domain.Education{}
	err := r.db.GetContext(ctx, edu, "SELECT * FROM education WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return edu, nil
}

func (r *educationRepository) Create(ctx context.Context, edu *domain.Education) error {
	query := `
		INSERT INTO education (institution, degree, start_date, end_date)
		VALUES (:institution, :degree, :start_date, :end_date)`
	_, err := r.db.NamedExecContext(ctx, query, edu)
	return err
}

func (r *educationRepository) Update(ctx context.Context, edu *domain.Education) error {
	query := `
		UPDATE education SET 
			institution = :institution, 
			degree = :degree, 
			start_date = :start_date, 
			end_date = :end_date
		WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, edu)
	return err
}

func (r *educationRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM education WHERE id = $1", id)
	return err
}
