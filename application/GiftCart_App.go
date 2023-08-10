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
	Create(c context.Context, giftCard *DTO.SendGiftCartRequest) error
	GetByID(c context.Context, id int) (*entity.GiftCard, error)
	GetByReceiverID(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error)
	GetBySenderID(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error)
	UpdateStatus(c context.Context, receiverID int, giftCartID int, status string) error
}

// Create isolating create from gift cart repo to use in interface layer
func (g *giftCartApp) Create(c context.Context, giftCard *DTO.SendGiftCartRequest) error {
	return g.gr.Create(c, giftCard)
}

// GetByID isolating getByID from gift cart repo to use in interface layer
func (g *giftCartApp) GetByID(c context.Context, id int) (*entity.GiftCard, error) {
	return g.gr.GetByID(c, id)
}

// GetByReceiverID isolating getByReceiverID from gift cart repo to use in interface layer
func (g *giftCartApp) GetByReceiverID(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error) {
	return g.gr.GetByReceiverID(c, receiverID, stat)
}

// GetBySenderID isolating getBySenderID from gift cart repo to use in interface layer
func (g *giftCartApp) GetBySenderID(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error) {
	return g.gr.GetBySenderID(c, senderID, stat)
}

// UpdateStatus isolating updateStatus from gift cart repo to use in interface layer
func (g *giftCartApp) UpdateStatus(c context.Context, receiverID int, giftCartID int, status string) error {
	return g.gr.UpdateStatus(c, receiverID, giftCartID, status)
}
