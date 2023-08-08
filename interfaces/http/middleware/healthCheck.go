package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"remote-task/infrastructure/persistence/mysql"
)

func HealthCheck(mysqlRepo *mysql.Repositories) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			isDown := false
			if err := mysqlRepo.Ping(); err != nil {
				log.Printf("Err http health check middleware, can't ping mysql %v\n", err)
				isDown = true
			}
			if isDown {
				return c.NoContent(http.StatusServiceUnavailable)
			}
			return next(c)
		}
	}
}
