package application

import (
	"context"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/entity"
	"remote-task/domain/giftCart/param"
	"remote-task/domain/giftCart/repository"
)

type giftCartApp struct {
	gr repository.GiftCardRepository
}

var _ GiftCartAppInterface = &giftCartApp{}

type GiftCartAppInterface interface {
	CreateGiftCard(c context.Context, giftCard *DTO.SendGiftCartRequest) error
	GetGiftCardByID(c context.Context, id int) (*entity.GiftCard, error)
	GetGiftCardsByReceiverID(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error)
	GetGiftCardsBySenderID(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error)
	GetGiftCardsByStatus(c context.Context, status string) ([]entity.GiftCard, error)
	UpdateGiftCardStatus(c context.Context, receiverID int, giftCartID int, status string) error
}

func (g *giftCartApp) CreateGiftCard(c context.Context, giftCard *DTO.SendGiftCartRequest) error {
	return g.gr.CreateGiftCard(c, giftCard)
}

func (g *giftCartApp) GetGiftCardByID(c context.Context, id int) (*entity.GiftCard, error) {
	return g.gr.GetGiftCardByID(c, id)
}

func (g *giftCartApp) GetGiftCardsByReceiverID(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error) {
	return g.gr.GetGiftCardsByReceiverID(c, receiverID, stat)
}

func (g *giftCartApp) GetGiftCardsBySenderID(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error) {
	return g.gr.GetGiftCardsBySenderID(c, senderID, stat)
}

func (g *giftCartApp) GetGiftCardsByStatus(c context.Context, status string) ([]entity.GiftCard, error) {
	return g.gr.GetGiftCardsByStatus(c, status)
}

func (g *giftCartApp) UpdateGiftCardStatus(c context.Context, receiverID int, giftCartID int, status string) error {
	return g.gr.UpdateGiftCardStatus(c, receiverID, giftCartID, status)
}
