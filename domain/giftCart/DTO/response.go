package DTO

import (
	"remote-task/domain/giftCart/param"
)

type SendGiftCartResponse struct {
	Message interface{} `json:"message"`
	Error   string      `json:"error"`
	Code    int         `json:"code"`
}

type UpdateGiftCartResponse struct {
	Message interface{} `json:"message"`
	Error   string      `json:"error"`
	Code    int         `json:"code"`
}

type GetMyGiftCartsResponse struct {
	Message Result `json:"message"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}
type Result struct {
	Count int                                `json:"count"`
	Data  []param.GiftCardJoinUserByReceiver `json:"'data'"`
}

type GetMySentGiftCartsResponse struct {
	Message SentResult `json:"message"`
	Error   string     `json:"error"`
	Code    int        `json:"code"`
}

type SentResult struct {
	Count int                              `json:"count"`
	Data  []param.GiftCardJoinUserBySender `json:"'data'"`
}
