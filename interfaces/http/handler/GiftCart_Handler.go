package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"remote-task/application"
	"remote-task/domain/giftCart/DTO"
)

type GiftCart struct {
	giftCartApp application.GiftCartAppInterface
	userApp     application.UserAppInterface
}

//GiftCart constructor
func NewGiftCartHandler(gcApp application.GiftCartAppInterface, uApp application.UserAppInterface) *GiftCart {
	return &GiftCart{
		giftCartApp: gcApp,
		userApp:     uApp,
	}
}

func (gc *GiftCart) SendGiftCart(c echo.Context) error {
	var req DTO.SendGiftCartRequest
	var res DTO.SendGiftCartResponse
	res.Code = http.StatusCreated
	res.Error = ""
	res.Message = "created"
	if err := c.Bind(&req); err != nil {
		res.Code = http.StatusBadRequest
		res.Error = err.Error()
		res.Message = "validation error"
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err2 := gc.giftCartApp.CreateGiftCard(c.Request().Context(), &req)
	if err2 != nil {
		res.Code = http.StatusBadRequest
		res.Error = err2.Error()
		res.Message = "validation error"
		return c.JSON(res.Code, res)
	}
	return c.JSON(http.StatusOK, res)
}
