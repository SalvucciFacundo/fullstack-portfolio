package repositories

import (
	"context"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/jmoiron/sqlx"
)

type socialRepository struct {
	db *sqlx.DB
}

func NewSocialRepository(db *sqlx.DB) ports.SocialRepository {
	return &socialRepository{db: db}
}

func (r *socialRepository) GetAll(ctx context.Context) ([]domain.SocialLink, error) {
	var links []domain.SocialLink
	err := r.db.SelectContext(ctx, &links, "SELECT * FROM social_links WHERE is_active = true")
	return links, err
}

func (r *socialRepository) Update(ctx context.Context, social *domain.SocialLink) error {
	query := `
		UPDATE social_links SET 
			platform = :platform, 
			url = :url, 
			icon_name = :icon_name, 
			is_active = :is_active
		WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, social)
	return err
}
