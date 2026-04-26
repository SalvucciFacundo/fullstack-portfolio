package ports

import (
	"context"
	"portfolio-backend/internal/core/domain"
)

// UserRepository defines the contract for user persistence.
type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
}
