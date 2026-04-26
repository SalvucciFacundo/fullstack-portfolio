package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"portfolio-backend/internal/core/domain"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProjectService is a mock implementation of ports.ProjectService
type MockProjectService struct {
	mock.Mock
}

func (m *MockProjectService) GetAll(ctx context.Context) ([]domain.Project, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Project), args.Error(1)
}

func (m *MockProjectService) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Project), args.Error(1)
}

func (m *MockProjectService) Create(ctx context.Context, p *domain.Project) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockProjectService) Update(ctx context.Context, p *domain.Project) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockProjectService) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProjectHandler_GetAll(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockProjectService)
	expectedProjects := []domain.Project{
		{Title: "Test Project 1", Description: "Description 1"},
	}
	mockService.On("GetAll", mock.Anything).Return(expectedProjects, nil)

	handler := NewProjectHandler(mockService)

	// Assertions
	if assert.NoError(t, handler.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var actualProjects []domain.Project
		err := json.Unmarshal(rec.Body.Bytes(), &actualProjects)
		assert.NoError(t, err)
		assert.Equal(t, len(expectedProjects), len(actualProjects))
		assert.Equal(t, expectedProjects[0].Title, actualProjects[0].Title)
	}
}
