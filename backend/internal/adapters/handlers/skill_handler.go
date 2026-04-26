package handlers

import (
	"net/http"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type skillHandler struct {
	service ports.SkillService
}

func NewSkillHandler(service ports.SkillService) *skillHandler {
	return &skillHandler{service: service}
}

func (h *skillHandler) GetAll(c echo.Context) error {
	data, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func (h *skillHandler) Create(c echo.Context) error {
	var s domain.Skill
	if err := c.Bind(&s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if err := h.service.Create(c.Request().Context(), &s); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, s)
}

func (h *skillHandler) Update(c echo.Context) error {
	var s domain.Skill
	if err := c.Bind(&s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if err := h.service.Update(c.Request().Context(), &s); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, s)
}

func (h *skillHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
