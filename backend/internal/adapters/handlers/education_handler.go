package handlers

import (
	"net/http"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type educationHandler struct {
	service ports.EducationService
}

func NewEducationHandler(service ports.EducationService) *educationHandler {
	return &educationHandler{service: service}
}

func (h *educationHandler) GetAll(c echo.Context) error {
	data, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func (h *educationHandler) Create(c echo.Context) error {
	var e domain.Education
	if err := c.Bind(&e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if err := h.service.Create(c.Request().Context(), &e); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, e)
}

func (h *educationHandler) Update(c echo.Context) error {
	var e domain.Education
	if err := c.Bind(&e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if err := h.service.Update(c.Request().Context(), &e); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, e)
}

func (h *educationHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
