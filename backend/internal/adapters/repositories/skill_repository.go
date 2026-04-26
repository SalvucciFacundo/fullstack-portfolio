package repositories

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type skillRepository struct {
	db *sqlx.DB
}

func NewSkillRepository(db *sqlx.DB) ports.SkillRepository {
	return &skillRepository{db: db}
}

func (r *skillRepository) GetAll(ctx context.Context) ([]domain.Skill, error) {
	var skills []domain.Skill
	err := r.db.SelectContext(ctx, &skills, "SELECT * FROM skills ORDER BY display_order ASC")
	return skills, err
}

func (r *skillRepository) GetByID(ctx context.Context, id string) (*domain.Skill, error) {
	skill := &domain.Skill{}
	err := r.db.GetContext(ctx, skill, "SELECT * FROM skills WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *skillRepository) Create(ctx context.Context, skill *domain.Skill) error {
	query := `
		INSERT INTO skills (name, icon_class, category, display_order)
		VALUES (:name, :icon_class, :category, :display_order)`
	_, err := r.db.NamedExecContext(ctx, query, skill)
	return err
}

func (r *skillRepository) Update(ctx context.Context, skill *domain.Skill) error {
	query := `
		UPDATE skills SET 
			name = :name, 
			icon_class = :icon_class, 
			category = :category, 
			display_order = :display_order
		WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, skill)
	return err
}

func (r *skillRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM skills WHERE id = $1", id)
	return err
}
