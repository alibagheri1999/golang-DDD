package test_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"remote-task/domain/giftCart/DTO"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// path	/gift-cart/send
func TestSentGiftCart(t *testing.T) {
	samples := []struct {
		inputJSON  []byte
		statusCode int
	}{
		{
			inputJSON:  []byte(`{"receiver_id": 0, "sender_id": 2,"amount": 10}`),
			statusCode: 422,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 1, "sender_id": 0,"amount": 0}`),
			statusCode: 422,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 0, "sender_id": 1,"amount": 0}`),
			statusCode: 422,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 1000, "sender_id": 1,"amount": 1}`),
			statusCode: 404,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 1, "sender_id": 1, "amount": 1}`),
			statusCode: 201,
		},
	}

	for _, v := range samples {
		bodyReader := bytes.NewReader(v.inputJSON)

		requestURL := fmt.Sprintf("http://localhost:%d/api/v1/gift-cart/send", 8080)
		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		client := http.Client{}
		res, err1 := client.Do(req)
		if err1 != nil {
			t.Fatal(err1)
		}
		assert.Equal(t, v.statusCode, res.StatusCode)
	}
}

// path	/gift-cart/update-status
func TestUpdateGiftCart(t *testing.T) {
	samples := []struct {
		inputJSON  []byte
		statusCode int
	}{
		{
			inputJSON:  []byte(`{"receiver_id": 0, "status": "0","gift_cart_id": 0}`),
			statusCode: 422,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 1, "status": "","gift_cart_id": 0}`),
			statusCode: 422,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 0, "status": "1","gift_cart_id": 0}`),
			statusCode: 422,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 2, "status": "accept","gift_cart_id": 1}`),
			statusCode: 200,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 2, "status": "reject", "gift_cart_id": 2}`),
			statusCode: 200,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 2, "status": "aaaa", "gift_cart_id": 3}`),
			statusCode: 400,
		},
		{
			inputJSON:  []byte(`{"receiver_id": 2, "status": "reject", "gift_cart_id": 1500000}`),
			statusCode: 404,
		},
	}

	for _, v := range samples {
		bodyReader := bytes.NewReader(v.inputJSON)
		requestURL := fmt.Sprintf("http://localhost:%d/api/v1/gift-cart/update-status", 8080)
		req, err := http.NewRequest(http.MethodPatch, requestURL, bodyReader)
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		client := http.Client{}
		res, err1 := client.Do(req)
		if err1 != nil {
			t.Fatal(err1)
		}
		assert.Equal(t, v.statusCode, res.StatusCode)
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
			statusCode: 404,
		},
	}

	for _, v := range samples {
		requestURL := fmt.Sprintf("http://localhost:%d/api/v1/gift-cart/my-carts/%s/%s", 8080, v.receiverID, v.types)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		client := http.Client{}
		res, err1 := client.Do(req)
		if err1 != nil {
			t.Fatal(err1)
		}
		result, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			t.Fatal(err2)
		}
		res.Body.Close()
		var response DTO.GetMyGiftCartsResponse
		err3 := json.Unmarshal([]byte(string(result)), &response)

		if err3 != nil {
			t.Errorf("cannot unmarshal response: %v\n", err)
		}
		assert.Equal(t, v.statusCode, res.StatusCode)
		for i := 0; i < len(response.Message.Data); i++ {
			el := response.Message.Data[i]
			ir, err := strconv.Atoi(v.receiverID)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, el.ReceiverID, ir)
		}
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
		requestURL := fmt.Sprintf("http://localhost:%d/api/v1/gift-cart/send-carts/%s?status=%s", 8080, v.senderID, v.types)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}
		client := http.Client{}
		res, err1 := client.Do(req)
		if err1 != nil {
			t.Fatal(err1)
		}
		result, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			t.Fatal(err2)
		}
		res.Body.Close()
		var response DTO.GetMySentGiftCartsResponse
		err3 := json.Unmarshal([]byte(string(result)), &response)

		if err3 != nil {
			t.Errorf("cannot unmarshal response: %v\n", err)
		}
		assert.Equal(t, v.statusCode, res.StatusCode)
		for i := 0; i < len(response.Message.Data); i++ {
			el := response.Message.Data[i]
			ir, err := strconv.Atoi(v.senderID)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, el.SenderID, ir)
		}
	}
}
