package repository

import (
	"context"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/entity"
	"remote-task/domain/giftCart/param"
)

type GiftCardRepository interface {
	Create(c context.Context, giftCard *DTO.SendGiftCartRequest) error
	GetByID(c context.Context, id int) (*entity.GiftCard, error)
	GetByReceiverID(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error)
	GetBySenderID(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error)
	UpdateStatus(c context.Context, receiverID int, giftCartID int, status string) error
}
