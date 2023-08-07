package routers

import "github.com/labstack/echo/v4"

func RegisterRoutes(router *echo.Echo, handler *handler.Handlers) {

	v1 := router.Group("/api/v1")

	v1.GET("/health-check", handler.HealthCheck)

	// gift cart api
	v1.POST("/gift-cart/send", handler.HealthCheck)           // sent cart with sent status
	v1.PATCH("/gift-cart/update-status", handler.HealthCheck) // update status
	v1.GET("/gift-cart/my-carts/:type", handler.HealthCheck)  // all of my carts by params if 1 is acc and if 2 is rej
	v1.GET("/gift-cart/send-carts", handler.HealthCheck)      // query status
	v1.GET("/gift-cart/pending-carts", handler.HealthCheck)   // get sent carts for updating them
}
