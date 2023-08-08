package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"remote-task/application"
	"remote-task/domain/giftCart/DTO"
)

type GiftCart struct {
	gca application.GiftCartAppInterface
	ua  application.UserAppInterface
}

//GiftCart constructor
func NewGiftCartHandler(gcApp application.GiftCartAppInterface, uApp application.UserAppInterface) *GiftCart {
	return &GiftCart{
		gca: gcApp,
		ua:  uApp,
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
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	sUser, err1 := gc.ua.GetUserByID(c.Request().Context(), req.ReceiverID)
	if err1 != nil || sUser.Username == "" {
		res.Code = http.StatusNotFound
		res.Error = err1.Error()
		res.Message = "Receiver error"
		return echo.NewHTTPError(http.StatusNotFound, res)
	}
	rUser, err2 := gc.ua.GetUserByID(c.Request().Context(), req.SenderID)
	if err1 != nil || rUser.Username == "" {
		res.Code = http.StatusNotFound
		res.Error = err2.Error()
		res.Message = "Sender error"
		return echo.NewHTTPError(http.StatusNotFound, res)
	}
	err3 := gc.gca.CreateGiftCard(c.Request().Context(), &req)
	if err2 != nil {
		res.Code = http.StatusBadRequest
		res.Error = err3.Error()
		res.Message = "validation error"
		return c.JSON(res.Code, res)
	}
	return c.JSON(http.StatusOK, res)
}
