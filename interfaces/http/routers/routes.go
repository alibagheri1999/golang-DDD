package routers

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/swag"
	"remote-task/interfaces/http/handler"
)

func RegisterRoutes(router *echo.Echo, handler *handler.Handlers) {

	v1 := router.Group("/api/v1")

	v1.GET("/health-check", handler.GeneralService.CheckHealth)

	// gift cart api
	v1.POST("/gift-cart/send", handler.GiftCartService.SendGiftCart)
	v1.PATCH("/gift-cart/update-status", handler.GiftCartService.UpdateGiftCart)
	v1.GET("/gift-cart/my-carts/:receiverID/:type", handler.GiftCartService.GetMyGiftCarts)
	v1.GET("/gift-cart/send-carts/:senderID", handler.GiftCartService.GetMySentCarts)

	// swagger
	v1.Any("/swagger/*", echoSwagger.WrapHandler)
}
