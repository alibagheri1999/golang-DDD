package repository

import (
	"context"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/entity"
)

type GiftCardRepository interface {
	CreateGiftCard(c context.Context, giftCard *DTO.SendGiftCartRequest) error
	GetGiftCardByID(c context.Context, id int) (*entity.GiftCard, error)
	GetGiftCardsByReceiverID(c context.Context, receiverID int, stat int) ([]entity.GiftCardJoinUserByReceiver, error)
	GetGiftCardsBySenderID(c context.Context, senderID int, stat int) ([]entity.GiftCardJoinUserBySender, error)
	GetGiftCardsByStatus(c context.Context, status string) ([]entity.GiftCard, error)
	UpdateGiftCardStatus(c context.Context, receiverID int, giftCartID int, status string) error
}
