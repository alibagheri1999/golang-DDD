package DTO

import "remote-task/domain/giftCart/entity"

type SendGiftCartResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}

type UpdateGiftCartResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}

type GetMyGiftCartsResponse struct {
	Message Result `json:"message"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}
type Result struct {
	Count int                                 `json:"count"`
	Data  []entity.GiftCardJoinUserByReceiver `json:"'data'"`
}

type GetMySentGiftCartsResponse struct {
	Message SentResult `json:"message"`
	Error   string     `json:"error"`
	Code    int        `json:"code"`
}

type SentResult struct {
	Count int                               `json:"count"`
	Data  []entity.GiftCardJoinUserBySender `json:"'data'"`
}
