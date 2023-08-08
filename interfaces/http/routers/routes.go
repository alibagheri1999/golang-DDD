package routers

import (
	"github.com/labstack/echo/v4"
	"remote-task/interfaces/http/handler"
)

func RegisterRoutes(router *echo.Echo, handler *handler.Handlers) {

	v1 := router.Group("/api/v1")

	v1.GET("/health-check", handler.GeneralService.CheckHealth)

	// gift cart api
	v1.POST("/gift-cart/send", handler.GiftCartService.SendGiftCart)                        // sent cart with sent status
	v1.PATCH("/gift-cart/update-status", handler.GiftCartService.UpdateGiftCart)            // update status
	v1.GET("/gift-cart/my-carts/:receiverID/:type", handler.GiftCartService.GetMyGiftCarts) // all of my carts by params if 1 is acc and if 2 is rej
	v1.GET("/gift-cart/send-carts/:senderID", handler.GiftCartService.GetMySentCarts)       // query status
}
