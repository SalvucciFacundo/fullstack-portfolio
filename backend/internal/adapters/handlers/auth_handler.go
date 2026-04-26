package handlers

import (
	"net/http"
	"portfolio-backend/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	service ports.AuthAppService
}

// NewAuthHandler creates a new instance of the authentication HTTP handler.
func NewAuthHandler(service ports.AuthAppService) *authHandler {
	return &authHandler{service: service}
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *authHandler) Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	token, err := h.service.Login(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
