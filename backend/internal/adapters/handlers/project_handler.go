package handlers

import (
	"net/http"
	"portfolio-backend/internal/core/domain"
	"portfolio-backend/internal/core/ports"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type projectHandler struct {
	service ports.ProjectService
}

// NewProjectHandler creates a new instance of the project HTTP handler.
func NewProjectHandler(service ports.ProjectService) *projectHandler {
	return &projectHandler{service: service}
}

func (h *projectHandler) GetAll(c echo.Context) error {
	projects, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, projects)
}

func (h *projectHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	project, err := h.service.GetByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "project not found")
	}
	return c.JSON(http.StatusOK, project)
}

func (h *projectHandler) Create(c echo.Context) error {
	var p domain.Project
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := h.service.Create(c.Request().Context(), &p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p)
}

func (h *projectHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var p domain.Project
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid project id")
	}
	p.ID = parsedID

	if err := h.service.Update(c.Request().Context(), &p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}

func (h *projectHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
