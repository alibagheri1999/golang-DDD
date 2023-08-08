package DTO

type SendGiftCartRequest struct {
	SenderID   int     `json:"sender_id"`
	ReceiverID int     `json:"receiver_id"`
	Amount     float64 `json:"amount"`
}

type UpdateGiftCartRequest struct {
	ReceiverID int    `json:"receiver_id"`
	Status     string `json:"status"`
	GiftCartID int    `json:"gift_cart_id"`
}
