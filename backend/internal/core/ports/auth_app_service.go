package ports

import "context"

// AuthAppService defines the application logic for authentication.
type AuthAppService interface {
	Login(ctx context.Context, email, password string) (string, error) // Returns JWT token
}
