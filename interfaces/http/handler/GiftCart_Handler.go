package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"remote-task/application"
	"remote-task/domain/giftCart/DTO"
	"strconv"
	"strings"
)

type GiftCart struct {
	gca application.GiftCartAppInterface
	ua  application.UserAppInterface
}

// GiftCart constructor
func NewGiftCartHandler(gcApp application.GiftCartAppInterface, uApp application.UserAppInterface) *GiftCart {
	return &GiftCart{
		gca: gcApp,
		ua:  uApp,
	}
}

// SendGiftCart  send Gift Cart godoc
// @Summary      sending gift cart
// @Description  sending gift cart from sender to receiver
// @Tags         GiftCart
// @Param        request  body  string  true  "request body" SchemaExample({"sender_id":"int","receiver_id":"int","amount": float64})
// @Success      201   {object}  DTO.SendGiftCartResponse
// @Failure      400   {object}  DTO.SendGiftCartResponse
// @Failure      404   {object}  DTO.SendGiftCartResponse
// @Router       /api/v1/gift-cart/send [post]
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
	if err2 != nil || rUser.Username == "" {
		res.Code = http.StatusNotFound
		res.Error = err2.Error()
		res.Message = "Sender error"
		return echo.NewHTTPError(http.StatusNotFound, res)
	}
	err3 := gc.gca.CreateGiftCard(c.Request().Context(), &req)
	if err3 != nil {
		res.Code = http.StatusBadRequest
		res.Error = err3.Error()
		res.Message = "validation error"
		return c.JSON(res.Code, res)
	}
	return c.JSON(http.StatusOK, res)
}

// UpdateGiftCart  update Gift Cart godoc
// @Summary      update gift cart
// @Description  update status of gift cart to accept or reject
// @Tags         GiftCart
// @Param        request  body  string  true  "request body" SchemaExample({"gift_cart_id":"int","receiver_id":"int","status": string})
// @Success      200   {object}  DTO.UpdateGiftCartResponse
// @Failure      400   {object}  DTO.UpdateGiftCartResponse
// @Failure      404   {object}  DTO.UpdateGiftCartResponse
// @Router       /api/v1/gift-cart/update-status [patch]
func (gc *GiftCart) UpdateGiftCart(c echo.Context) error {
	var req DTO.UpdateGiftCartRequest
	var res DTO.UpdateGiftCartResponse
	res.Code = http.StatusOK
	res.Error = ""
	res.Message = "updated"
	if err := c.Bind(&req); err != nil {
		res.Code = http.StatusBadRequest
		res.Error = err.Error()
		res.Message = "validation error"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	status := strings.ToLower(req.Status)
	if status != "accept" && status != "reject" {
		res.Code = http.StatusBadRequest
		res.Error = "status must be accept or reject"
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
	err2 := gc.gca.UpdateGiftCardStatus(c.Request().Context(), req.ReceiverID, req.GiftCartID, req.Status)
	if err2 != nil {
		res.Code = http.StatusBadRequest
		res.Error = err2.Error()
		res.Message = "validation error"
		return c.JSON(res.Code, res)
	}
	return c.JSON(http.StatusOK, res)
}

// GetMyGiftCarts  get Gift Cart godoc
// @Summary      get gift cart
// @Description  getting all gift cart if stat is 1 it is accepted and if it is 2 we get rejected gift cart and at the end with number 3 we will see all sent status gift carts
// @Tags         GiftCart
// @Param        type  path      string  true  "type param"  receiverID  path      string  true  "receiverID param"
// @Success      200   {object}  DTO.GetMyGiftCartsResponse
// @Failure      400   {object}  DTO.GetMyGiftCartsResponse
// @Failure      404   {object}  DTO.GetMyGiftCartsResponse
// @Router       /api/v1/gift-cart/my-carts/:receiverID/:type [get]
func (gc *GiftCart) GetMyGiftCarts(c echo.Context) error {
	var res DTO.GetMyGiftCartsResponse
	var result DTO.Result
	res.Code = http.StatusOK
	res.Error = ""
	res.Message = result
	stat := c.Param("type")
	ReceiverID := c.Param("receiverID")
	iStat, conErr1 := strconv.Atoi(stat)
	iReceiverID, conErr2 := strconv.Atoi(ReceiverID)
	if conErr1 != nil || conErr2 != nil {
		res.Code = http.StatusBadRequest
		res.Error = "converting failed"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	if iStat != 1 && iStat != 2 && iStat != 3 {
		res.Code = http.StatusBadRequest
		res.Error = "type must be 1 or 2"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	sUser, err2 := gc.ua.GetUserByID(c.Request().Context(), iReceiverID)
	if err2 != nil || sUser.Username == "" {
		res.Code = http.StatusNotFound
		res.Error = err2.Error()
		return echo.NewHTTPError(http.StatusNotFound, res)
	}
	r, err3 := gc.gca.GetGiftCardsByReceiverID(c.Request().Context(), iReceiverID, iStat)
	if err3 != nil {
		res.Code = http.StatusBadRequest
		res.Error = err3.Error()
		return c.JSON(res.Code, res)
	}
	result.Data = r
	result.Count = len(r)
	res.Message = result
	return c.JSON(http.StatusOK, res)
}

// GetMySentCarts  get sent Gift Cart godoc
// @Summary      get sent gift cart
// @Description  getting all sent gift cart if status is 1 it is accepted and if it is 2 we get rejected gift cart and at the end with number 3 we will see all sent status gift carts and anything expect these 3 number will give us all data
// @Tags         GiftCart
// @Param        senderID  path      string  true  "senderID param"  status  query  string  false  "status of gift"
// @Success      200   {object}  DTO.GetMySentGiftCartsResponse
// @Failure      400   {object}  DTO.GetMySentGiftCartsResponse
// @Failure      404   {object}  DTO.GetMySentGiftCartsResponse
// @Router       /api/v1/gift-cart/my-carts/:receiverID/:type [get]
func (gc *GiftCart) GetMySentCarts(c echo.Context) error {
	var res DTO.GetMySentGiftCartsResponse
	var result DTO.SentResult
	res.Code = http.StatusOK
	res.Error = ""
	res.Message = result
	SenderID := c.Param("senderID")
	iStatus, err1 := strconv.Atoi(c.QueryParam("status"))
	iSenderID, err1 := strconv.Atoi(SenderID)
	if err1 != nil {
		res.Code = http.StatusBadRequest
		res.Error = err1.Error()
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	sUser, err2 := gc.ua.GetUserByID(c.Request().Context(), iSenderID)
	if err2 != nil || sUser.Username == "" {
		res.Code = http.StatusNotFound
		res.Error = err2.Error()
		return echo.NewHTTPError(http.StatusNotFound, res)
	}
	r, err3 := gc.gca.GetGiftCardsBySenderID(c.Request().Context(), iSenderID, iStatus)
	if err3 != nil {
		res.Code = http.StatusBadRequest
		res.Error = err3.Error()
		return c.JSON(res.Code, res)
	}
	result.Data = r
	result.Count = len(r)
	res.Message = result
	return c.JSON(http.StatusOK, res)
}
