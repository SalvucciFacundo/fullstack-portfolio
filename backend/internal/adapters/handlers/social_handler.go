package handlers

import (
	"net/http"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type socialHandler struct {
	service ports.SocialService
}

func NewSocialHandler(service ports.SocialService) *socialHandler {
	return &socialHandler{service: service}
}

func (h *socialHandler) GetAll(c echo.Context) error {
	data, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func (h *socialHandler) Update(c echo.Context) error {
	var s domain.SocialLink
	if err := c.Bind(&s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if err := h.service.Update(c.Request().Context(), &s); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, s)
}
