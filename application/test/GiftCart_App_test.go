package test_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"remote-task/application"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/entity"
	"remote-task/domain/giftCart/param"
	"testing"
	"time"
)

var (
	CreateRepo          func(c context.Context, giftCard *DTO.SendGiftCartRequest) error
	GetByIDRepo         func(c context.Context, id int) (*entity.GiftCard, error)
	GetByReceiverIDRepo func(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error)
	GetBySenderIDRepo   func(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error)
	UpdateStatusRepo    func(c context.Context, receiverID int, giftCartID int, status string) error
)

type fakeGiftCartRepo struct{}

func (fg *fakeGiftCartRepo) Create(c context.Context, giftCard *DTO.SendGiftCartRequest) error {
	return CreateRepo(c, giftCard)
}

func (fg *fakeGiftCartRepo) GetByID(c context.Context, id int) (*entity.GiftCard, error) {
	return GetByIDRepo(c, id)
}

func (fg *fakeGiftCartRepo) GetByReceiverID(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error) {
	return GetByReceiverIDRepo(c, receiverID, stat)
}

func (fg *fakeGiftCartRepo) GetBySenderID(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error) {
	return GetBySenderIDRepo(c, senderID, stat)
}

func (fg *fakeGiftCartRepo) UpdateStatus(c context.Context, receiverID int, giftCartID int, status string) error {
	return UpdateStatusRepo(c, receiverID, giftCartID, status)
}

var giftCartAppFake application.GiftCartAppInterface = &fakeGiftCartRepo{}

func TestCreate_Success(t *testing.T) {
	c := context.Background()

	CreateRepo = func(c context.Context, giftCard *DTO.SendGiftCartRequest) error {
		if giftCard.ReceiverID == 0 || giftCard.SenderID == 0 {
			return errors.New("must be bigger than 0")
		}
		return nil
	}
	giftCart := &DTO.SendGiftCartRequest{
		SenderID:   1,
		ReceiverID: 2,
		Amount:     0.99,
	}
	err := giftCartAppFake.Create(c, giftCart)
	assert.Nil(t, err)
}

func TestGiftCartGetByID_Success(t *testing.T) {
	c := context.Background()

	GetByIDRepo = func(c context.Context, id int) (*entity.GiftCard, error) {
		return &entity.GiftCard{
			ID:         1,
			SenderID:   2,
			ReceiverID: 3,
			Amount:     3,
			Status:     "sent",
			CreatedAt:  time.Now(),
		}, nil
	}
	r, err := giftCartAppFake.GetByID(c, 1)
	assert.Nil(t, err)
	assert.EqualValues(t, r.SenderID, 2)
	assert.EqualValues(t, r.ReceiverID, 3)
	assert.EqualValues(t, r.Amount, 3)
	assert.EqualValues(t, r.Status, "sent")
}

func TestGetByReceiverID_Success(t *testing.T) {
	c := context.Background()

	GetByReceiverIDRepo = func(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error) {
		var status string
		switch stat {
		case 1:
			status = "accept"
		case 2:
			status = "reject"
		case 3:
			status = "sent"
		}
		myResult := param.GiftCardJoinUserByReceiver{
			GiftCardID:        1,
			SenderID:          2,
			ReceiverID:        3,
			Amount:            1.99,
			Status:            status,
			GiftCardCreatedAt: time.Now(),
			ReceiverName:      "string",
			ReceiverEmail:     "string",
			UserID:            3,
		}
		myResultArr := make([]param.GiftCardJoinUserByReceiver, 0)
		myResultArr = append(myResultArr, myResult)
		return myResultArr, nil
	}
	for i := 1; i < 4; i++ {
		var status string
		switch i {
		case 1:
			status = "accept"
		case 2:
			status = "reject"
		case 3:
			status = "sent"
		}
		r, err := giftCartAppFake.GetByReceiverID(c, 3, i)
		assert.Nil(t, err)
		assert.EqualValues(t, r[0].SenderID, 2)
		assert.EqualValues(t, r[0].ReceiverID, 3)
		assert.EqualValues(t, r[0].Amount, 1.99)
		assert.EqualValues(t, r[0].Status, status)
		assert.EqualValues(t, r[0].ReceiverName, "string")
		assert.EqualValues(t, r[0].ReceiverEmail, "string")
	}
}

func TestGetBySenderID_Success(t *testing.T) {
	c := context.Background()

	GetBySenderIDRepo = func(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error) {
		var status string
		switch stat {
		case 1:
			status = "accept"
		case 2:
			status = "reject"
		case 3:
			status = "sent"
		}
		myResult := param.GiftCardJoinUserBySender{
			GiftCardID:        1,
			SenderID:          2,
			ReceiverID:        3,
			Amount:            1.99,
			Status:            status,
			GiftCardCreatedAt: time.Now(),
			SenderName:        "string",
			SenderEmail:       "string",
			UserID:            2,
		}
		myResultArr := make([]param.GiftCardJoinUserBySender, 0)
		myResultArr = append(myResultArr, myResult)
		return myResultArr, nil
	}
	for i := 1; i < 4; i++ {
		var status string
		switch i {
		case 1:
			status = "accept"
		case 2:
			status = "reject"
		case 3:
			status = "sent"
		}
		r, err := giftCartAppFake.GetBySenderID(c, 3, i)
		assert.Nil(t, err)
		assert.EqualValues(t, r[0].SenderID, 2)
		assert.EqualValues(t, r[0].ReceiverID, 3)
		assert.EqualValues(t, r[0].Amount, 1.99)
		assert.EqualValues(t, r[0].Status, status)
		assert.EqualValues(t, r[0].SenderEmail, "string")
		assert.EqualValues(t, r[0].SenderName, "string")
	}
}

func TestUpdate_Success(t *testing.T) {
	c := context.Background()

	UpdateStatusRepo = func(c context.Context, receiverID int, giftCartID int, status string) error {
		if giftCartID == 0 || receiverID == 0 {
			t.Fatal("must be bigger than 0")
		}
		if status != "accept" && status != "reject" {
			t.Fatal("must be accept or reject")
		}
		return nil
	}
	err := giftCartAppFake.UpdateStatus(c, 1, 1, "accept")
	assert.Nil(t, err)
}
