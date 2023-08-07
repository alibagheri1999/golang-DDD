package repository

import "remote-task/domain/giftCart/entity"

type GiftCardRepository interface {
	CreateGiftCard(giftCard *entity.GiftCard) error
	GetGiftCardByID(id int) (*entity.GiftCard, error)
	GetGiftCardsByReceiverID(receiverID int) ([]entity.GiftCardJoinUserByReceiver, error)
	GetGiftCardsBySenderID(senderID int) ([]entity.GiftCardJoinUserBySender, error)
	GetGiftCardsByStatus(status string) ([]entity.GiftCard, error)
	UpdateGiftCardStatus(id int, status string) error
}
