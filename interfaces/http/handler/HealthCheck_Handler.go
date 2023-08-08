package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type General struct{}

// General constructor
func NewGeneralHandler() *General {
	return &General{}
}

// CheckHealth HealthCheck godoc
// @Summary      Health check
// @Description  Health check
// @Tags         Health check
// @Success      204
// @Failure      400  "bad request"
// @Router       /api/v1/health-check [get]
func (g *General) CheckHealth(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
