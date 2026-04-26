package handlers

import (
	"net/http"
	"portfolio-backend/internal/core/ports"
	"strings"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware handles JWT validation for protected routes.
func AuthMiddleware(authService ports.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format")
			}

			token := parts[1]
			email, role, err := authService.ValidateToken(c.Request().Context(), token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			// Store user info in context for later use if needed
			c.Set("user_email", email)
			c.Set("user_role", role)

			return next(c)
		}
	}
}
