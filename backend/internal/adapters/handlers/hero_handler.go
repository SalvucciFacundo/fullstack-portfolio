package handlers

import (
	"net/http"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type heroHandler struct {
	service ports.HeroService
}

func NewHeroHandler(service ports.HeroService) *heroHandler {
	return &heroHandler{service: service}
}

func (h *heroHandler) Get(c echo.Context) error {
	data, err := h.service.Get(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func (h *heroHandler) Update(c echo.Context) error {
	var p domain.HeroSection
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if err := h.service.Update(c.Request().Context(), &p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}
