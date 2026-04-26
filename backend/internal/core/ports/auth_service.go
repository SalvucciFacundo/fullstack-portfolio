package ports

import "context"

// AuthService defines the contract for authentication logic.
type AuthService interface {
	GenerateToken(ctx context.Context, email, role string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, string, error) // Returns email, role, error
	HashPassword(password string) (string, error)
	ComparePassword(hash, password string) error
}
