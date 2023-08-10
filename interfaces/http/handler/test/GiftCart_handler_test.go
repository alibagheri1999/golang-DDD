package test_test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"remote-task/domain/giftCart/DTO"
	"testing"
)

// path	/gift-cart/send
func TestSentGiftCart(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"receiver_id": 0, "sender_id": 0,"amount": 0}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"receiver_id": 1, "sender_id": 0,"amount": 0}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"receiver_id": 0, "sender_id": 1,"amount": 0}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"receiver_id": 1000, "sender_id": 1,"amount": 1}`,
			statusCode: 404,
		},
		{
			inputJSON:  `{"receiver_id": 1, "sender_id": 1, "amount": 1}`,
			statusCode: 201,
		},
	}

	for _, v := range samples {
		r := echo.New()
		r.POST("/api/v1/gift-cart/send", g.SendGiftCart)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/gift-cart/send", bytes.NewBufferString(v.inputJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		log.Println(rr, v.statusCode)
	}
}

// path	/gift-cart/update-status
func TestUpdateGiftCart(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			inputJSON:  `{"receiver_id": 0, "status": "0","gift_cart_id": 0}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"receiver_id": 1, "status": "","gift_cart_id": 0}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"receiver_id": 0, "status": "1","gift_cart_id": 0}`,
			statusCode: 422,
		},
		{
			inputJSON:  `{"receiver_id": 1, "status": "accept","gift_cart_id": 2}`,
			statusCode: 200,
		},
		{
			inputJSON:  `{"receiver_id": 1, "status": "reject", "gift_cart_id": 2}`,
			statusCode: 200,
		},
		{
			inputJSON:  `{"receiver_id": 1, "status": "aaaa", "gift_cart_id": 1}`,
			statusCode: 400,
		},
	}

	for _, v := range samples {
		r := echo.New()
		r.POST("/api/v1/gift-cart/update-status", g.UpdateGiftCart)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/gift-cart/update-status", bytes.NewBufferString(v.inputJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		log.Println(rr, v.statusCode)
		assert.Equal(t, rr.Code, v.statusCode)
	}
}

// path	/gift-cart/my-carts/:receiverID/:type
func TestMyGiftCart(t *testing.T) {
	samples := []struct {
		receiverID string
		types      string
		statusCode int
	}{
		{
			receiverID: "1",
			types:      "1",
			statusCode: 200,
		},
		{
			receiverID: "1",
			types:      "2",
			statusCode: 200,
		},
		{
			receiverID: "1",
			types:      "3",
			statusCode: 200,
		},
		{
			receiverID: "1",
			types:      "4",
			statusCode: 400,
		},
		{
			receiverID: "1",
			types:      "400",
			statusCode: 400,
		},
		{
			receiverID: "100000",
			types:      "1",
			statusCode: 494,
		},
	}

	for _, v := range samples {
		r := echo.New()
		var requestBody bytes.Buffer
		r.GET("/api/v1/gift-cart/my-carts/:receiverID/:type", g.GetMyGiftCarts)
		req := httptest.NewRequest(http.MethodGet, `/api/v1/gift-cart/my-carts/`+v.receiverID+"/"+v.types, &requestBody)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		log.Println(rr, v.statusCode)
		var result []DTO.GetMyGiftCartsResponse
		err := json.Unmarshal(rr.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("cannot unmarshal response: %v\n", err)
		}
		for i := 0; i < len(result); i++ {
			el := result[i].Message.Data
			for j := 0; i < len(el); i++ {
				receiverID := el[j]
				assert.Equal(t, receiverID, v.receiverID)
			}

		}
		assert.Equal(t, rr.Code, v.statusCode)
	}
}

// path	/gift-cart/send-carts/:senderID
func TestMySendGiftCart(t *testing.T) {
	samples := []struct {
		senderID   string
		types      string
		statusCode int
	}{
		{
			senderID:   "1",
			types:      "1",
			statusCode: 200,
		},
		{
			senderID:   "1",
			types:      "2",
			statusCode: 200,
		},
		{
			senderID:   "1",
			types:      "3",
			statusCode: 200,
		},
		{
			senderID:   "1",
			types:      "4",
			statusCode: 200,
		},
		{
			senderID:   "10000",
			types:      "1",
			statusCode: 404,
		},
	}

	for _, v := range samples {
		r := echo.New()
		var requestBody bytes.Buffer
		r.GET("/api/v1/gift-cart/send-carts/:senderID", g.GetMySentCarts)
		req := httptest.NewRequest(http.MethodGet, `/api/v1/gift-cart/send-carts/`+v.senderID+"?status="+v.types, &requestBody)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		log.Println(rr, v.statusCode)
		var result []DTO.GetMySentGiftCartsResponse
		err := json.Unmarshal(rr.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("cannot unmarshal response: %v\n", err)
		}
		for i := 0; i < len(result); i++ {
			el := result[i].Message.Data
			for j := 0; i < len(el); i++ {
				receiverID := el[j]
				assert.Equal(t, receiverID, v.senderID)
			}

		}
		assert.Equal(t, rr.Code, v.statusCode)
	}
}
