package application

import (
	"context"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/entity"
	"remote-task/infrastructure/persistence/mysql"
	"time"
)

type giftCartApp struct {
	gr mysql.GiftCardRepositoryImpl
}

var _ GiftCartAppInterface = &giftCartApp{}

type GiftCartAppInterface interface {
	CreateGiftCard(c context.Context, giftCard *DTO.SendGiftCartRequest) error
	GetGiftCardByID(c context.Context, id int) (*entity.GiftCard, error)
	GetGiftCardsByReceiverID(c context.Context, receiverID int) ([]entity.GiftCardJoinUserByReceiver, error)
	GetGiftCardsBySenderID(c context.Context, senderID int) ([]entity.GiftCardJoinUserBySender, error)
	GetGiftCardsByStatus(c context.Context, status string) ([]entity.GiftCard, error)
	UpdateGiftCardStatus(c context.Context, id int, status string) error
}

func (g *giftCartApp) CreateGiftCard(c context.Context, giftCard *DTO.SendGiftCartRequest) error {
	gc := entity.GiftCard{
		CreatedAt:  time.Now(),
		Status:     "sent",
		Amount:     giftCard.Amount,
		SenderID:   giftCard.SenderID,
		ReceiverID: giftCard.ReceiverID,
	}
	return g.gr.CreateGiftCard(c, &gc)
}

func (g *giftCartApp) GetGiftCardByID(c context.Context, id int) (*entity.GiftCard, error) {
	return g.gr.GetGiftCardByID(c, id)
}

func (g *giftCartApp) GetGiftCardsByReceiverID(c context.Context, receiverID int) ([]entity.GiftCardJoinUserByReceiver, error) {
	return g.gr.GetGiftCardsByReceiverID(c, receiverID)
}

func (g *giftCartApp) GetGiftCardsBySenderID(c context.Context, senderID int) ([]entity.GiftCardJoinUserBySender, error) {
	return g.gr.GetGiftCardsBySenderID(c, senderID)
}

func (g *giftCartApp) GetGiftCardsByStatus(c context.Context, status string) ([]entity.GiftCard, error) {
	return g.gr.GetGiftCardsByStatus(c, status)
}

func (g *giftCartApp) UpdateGiftCardStatus(c context.Context, id int, status string) error {
	return g.gr.UpdateGiftCardStatus(c, id, status)
}
