package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorResponse represents the structured error format for the frontend.
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// CustomHTTPErrorHandler handles errors globally and returns a structured JSON response.
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if m, ok := he.Message.(string); ok {
			message = m
		}
	}

	// For the frontend Toasts, we always want this structure
	resp := ErrorResponse{
		Status:  "error",
		Message: message,
		Code:    code,
	}

	// Log the error if it's a 500
	if code >= 500 {
		c.Logger().Error(err)
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, resp)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
