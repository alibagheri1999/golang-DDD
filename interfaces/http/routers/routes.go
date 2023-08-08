package routers

import (
	"github.com/labstack/echo/v4"
	"remote-task/interfaces/http/handler"
)

func RegisterRoutes(router *echo.Echo, handler *handler.Handlers) {

	v1 := router.Group("/api/v1")

	v1.GET("/health-check", handler.GiftCartService.SendGiftCart)

	// gift cart api
	v1.POST("/gift-cart/send", handler.GiftCartService.SendGiftCart)           // sent cart with sent status
	v1.PATCH("/gift-cart/update-status", handler.GiftCartService.SendGiftCart) // update status
	v1.GET("/gift-cart/my-carts/:type", handler.GiftCartService.SendGiftCart)  // all of my carts by params if 1 is acc and if 2 is rej
	v1.GET("/gift-cart/send-carts", handler.GiftCartService.SendGiftCart)      // query status
	v1.GET("/gift-cart/pending-carts", handler.GiftCartService.SendGiftCart)   // get sent carts for updating them
}
