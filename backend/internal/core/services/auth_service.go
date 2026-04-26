package services

import (
	"context"
	"errors"
	"portfolio-backend/internal/core/ports"
)

type authAppService struct {
	userRepo    ports.UserRepository
	authAdapter ports.AuthService
}

// NewAuthAppService creates a new instance of the authentication application service.
func NewAuthAppService(userRepo ports.UserRepository, authAdapter ports.AuthService) *authAppService {
	return &authAppService{
		userRepo:    userRepo,
		authAdapter: authAdapter,
	}
}

func (s *authAppService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = s.authAdapter.ComparePassword(user.PasswordHash, password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := s.authAdapter.GenerateToken(ctx, user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
