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

// SendGiftValidate validating the requirement fields
func (u *SendGiftCartRequest) SendGiftValidate() map[string]string {
	var errorMessages = make(map[string]string)
	if u.SenderID == 0 {
		errorMessages["sender_id_required"] = "SenderID required"
	}
	if u.ReceiverID == 0 {
		errorMessages["receiver_id_required"] = "ReceiverID required"
	}
	if u.Amount == 0 {
		errorMessages["amount_required"] = "Amount required"
	}
	return errorMessages
}

// UpdateGiftValidate validating the requirement fields
func (u *UpdateGiftCartRequest) UpdateGiftValidate() map[string]string {
	var errorMessages = make(map[string]string)
	if u.ReceiverID == 0 {
		errorMessages["receiver_id_required"] = "ReceiverID required"
	}
	if u.Status == "" {
		errorMessages["status_required"] = "status required"
	}
	if u.GiftCartID == 0 {
		errorMessages["gift_cart_id_required"] = "GiftCartID required"
	}
	return errorMessages
}
