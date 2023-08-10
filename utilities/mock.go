package utilities

import (
	"context"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/entity"
	"remote-task/domain/giftCart/param"
	userEntity "remote-task/domain/user/entity"
)

type GiftCartAppInterface struct {
	CreateRepo          func(c context.Context, giftCard *DTO.SendGiftCartRequest) error
	GetByIDRepo         func(c context.Context, id int) (*entity.GiftCard, error)
	GetByReceiverIDRepo func(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error)
	GetBySenderIDRepo   func(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error)
	UpdateStatusRepo    func(c context.Context, receiverID int, giftCartID int, status string) error
}

func (gai *GiftCartAppInterface) Create(c context.Context, giftCard *DTO.SendGiftCartRequest) error {
	return gai.CreateRepo(c, giftCard)
}

func (gai *GiftCartAppInterface) GetByID(c context.Context, id int) (*entity.GiftCard, error) {
	return gai.GetByIDRepo(c, id)
}

func (gai *GiftCartAppInterface) GetByReceiverID(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error) {
	return gai.GetByReceiverIDRepo(c, receiverID, stat)
}

func (gai *GiftCartAppInterface) GetBySenderID(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error) {
	return gai.GetBySenderIDRepo(c, senderID, stat)
}

func (gai *GiftCartAppInterface) UpdateStatus(c context.Context, receiverID int, giftCartID int, status string) error {
	return gai.UpdateStatusRepo(c, receiverID, giftCartID, status)
}

type UserAppInterface struct {
	GetUserByIDfn func(c context.Context, id int) (*userEntity.User, error)
}

func (gai *UserAppInterface) GetByID(c context.Context, id int) (*userEntity.User, error) {
	return gai.GetUserByIDfn(c, id)
}
