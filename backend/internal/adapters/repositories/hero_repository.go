package repositories

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type heroRepository struct {
	db *sqlx.DB
}

func NewHeroRepository(db *sqlx.DB) ports.HeroRepository {
	return &heroRepository{db: db}
}

func (r *heroRepository) Get(ctx context.Context) (*domain.HeroSection, error) {
	hero := &domain.HeroSection{}
	err := r.db.GetContext(ctx, hero, "SELECT * FROM hero_section LIMIT 1")
	if err != nil {
		return nil, err
	}
	return hero, nil
}

func (r *heroRepository) Update(ctx context.Context, hero *domain.HeroSection) error {
	query := `
		UPDATE hero_section SET 
			headline = :headline, 
			subheadline = :subheadline, 
			biography = :biography, 
			profile_image = :profile_image, 
			resume_url = :resume_url,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = :id`
	
	_, err := r.db.NamedExecContext(ctx, query, hero)
	return err
}
