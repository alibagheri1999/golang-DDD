package param

import "time"

type GiftCardJoinUserBySender struct {
	GiftCardID        int       `json:"id"`
	SenderID          int       `json:"sender_id"`
	ReceiverID        int       `json:"receiver_id"`
	Amount            float64   `json:"amount"`
	Status            string    `json:"status"`
	GiftCardCreatedAt time.Time `json:"created_at"`
	SenderName        string    `json:"username"`
	SenderEmail       string    `json:"email"`
	UserID            int       `json:"userID"`
}

type GiftCardJoinUserByReceiver struct {
	GiftCardID        int       `json:"id"`
	SenderID          int       `json:"sender_id"`
	ReceiverID        int       `json:"receiver_id"`
	Amount            float64   `json:"amount"`
	Status            string    `json:"status"`
	GiftCardCreatedAt time.Time `json:"created_at"`
	ReceiverName      string    `json:"username"`
	ReceiverEmail     string    `json:"email"`
	UserID            int       `json:"userID"`
}
