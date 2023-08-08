package repository

import (
	"context"
	"remote-task/domain/giftCart/entity"
)

type GiftCardRepository interface {
	CreateGiftCard(c context.Context, giftCard *entity.GiftCard) error
	GetGiftCardByID(c context.Context, id int) (*entity.GiftCard, error)
	GetGiftCardsByReceiverID(c context.Context, receiverID int) ([]entity.GiftCardJoinUserByReceiver, error)
	GetGiftCardsBySenderID(c context.Context, senderID int) ([]entity.GiftCardJoinUserBySender, error)
	GetGiftCardsByStatus(c context.Context, status string) ([]entity.GiftCard, error)
	UpdateGiftCardStatus(c context.Context, id int, status string) error
}
