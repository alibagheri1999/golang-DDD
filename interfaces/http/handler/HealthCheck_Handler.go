package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type General struct{}

//General constructor
func NewGeneralHandler() *General {
	return &General{}
}

func (g *General) CheckHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, "HEALTH_CHECK")
}
