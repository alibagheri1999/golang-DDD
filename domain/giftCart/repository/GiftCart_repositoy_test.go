package repository_test

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	appConfig "remote-task/config"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/giftCartConst"
	"remote-task/domain/giftCart/repository"
	"remote-task/infrastructure/persistence/mysql"
	"testing"
)

// Test_Gift_Cart_Repo test all methods in gift cart repository
func Test_Gift_Cart_Repo(t *testing.T) {
	appConfig.Init()
	dbCfg := appConfig.Get().Mysql
	repo, err := mysql.NewRepositories(dbCfg.Username, dbCfg.Password, dbCfg.Port, dbCfg.Host, dbCfg.Name)
	if err != nil {
		t.Fatal(err)
	}
	GiftCarRepo := repository.NewGiftCardRepository(repo)
	sentTestCase := []SentTestCase{
		{
			name:        "created successfully",
			ReceiverID:  4,
			SenderID:    3,
			Amount:      1.99,
			errExpected: nil,
		},
	}
	for _, tc := range sentTestCase {
		t.Run(tc.name, func(t *testing.T) {
			giftCart := &DTO.SendGiftCartRequest{
				SenderID:   tc.SenderID,
				ReceiverID: tc.ReceiverID,
				Amount:     tc.Amount,
			}
			err := GiftCarRepo.Create(context.Background(), giftCart)
			if !errors.Is(err, tc.errExpected) {
				t.Errorf("expected error %v, got %v", tc.errExpected, err)
			}
		})
	}
	getByReceiverIDTestCase := []GetByReceiverIDTestCase{
		{
			name:       "get accepted gift carts for user 1",
			ReceiverID: 1,
			stat:       1,
		},
		{
			name:       "get rejected gift carts for user 1",
			ReceiverID: 1,
			stat:       2,
		},
		{
			name:       "get sent gift carts for user 1",
			ReceiverID: 1,
			stat:       3,
		},
	}
	for _, tc := range getByReceiverIDTestCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err2 := GiftCarRepo.GetByReceiverID(context.Background(), tc.ReceiverID, tc.stat)
			if err2 != nil {
				t.Fatal(err2)
			}
			var status string
			switch tc.stat {
			case 1:
				status = "accept"
			case 2:
				status = "reject"
			case 3:
				status = "sent"
			}
			for i := 0; i < len(res); i++ {
				resStatus := res[i].Status
				if resStatus != status {
					t.Fatal("wrong result")
				}
			}
		})
	}
	getBySenderIDTestCase := []GetBySenderIDTestCase{
		{
			name:     "get accepted gift carts for user 2",
			SenderID: 2,
			stat:     1,
		},
		{
			name:     "get rejected gift carts for user 2",
			SenderID: 2,
			stat:     2,
		},
		{
			name:     "get sent gift carts for user 2",
			SenderID: 2,
			stat:     3,
		},
		{
			name:     "get all gift carts for user 2",
			SenderID: 2,
			stat:     4,
		},
		{
			name:     "get all gift carts for user 2",
			SenderID: 2,
			stat:     100,
		},
	}
	for _, tc := range getBySenderIDTestCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err2 := GiftCarRepo.GetBySenderID(context.Background(), tc.SenderID, tc.stat)
			if err2 != nil {
				t.Fatal(err2)
			}
			var status string
			switch tc.stat {
			case 1:
				status = "accept"
			case 2:
				status = "reject"
			case 3:
				status = "sent"
			default:
				status = "all"
			}
			for i := 0; i < len(res); i++ {
				resStatus := res[i].Status
				if status != "all" {
					if resStatus != status {
						t.Fatal("wrong result")
					}
				}
			}
		})
	}
	updateTestCase := []UpdateTestCase{
		{
			name:        "updated to accept",
			ReceiverID:  4,
			GiftCartID:  7,
			status:      "accept",
			errExpected: nil,
		},
		{
			name:        "updated to reject",
			ReceiverID:  4,
			GiftCartID:  8,
			status:      "reject",
			errExpected: nil,
		},
		{
			name:        "must get error because the status must be accept or reject",
			ReceiverID:  4,
			GiftCartID:  7,
			status:      "AAAA",
			errExpected: giftCartConst.ERR_UPDATE_GIFT_CART,
		},
		{
			name:        "must get error because GiftCartID does not exist",
			ReceiverID:  3,
			GiftCartID:  70,
			status:      "reject",
			errExpected: giftCartConst.ERR_NOT_FOUND,
		},
	}
	for _, tc := range updateTestCase {
		t.Run(tc.name, func(t *testing.T) {
			err = GiftCarRepo.UpdateStatus(context.Background(), tc.ReceiverID, tc.GiftCartID, tc.status)
			if !errors.Is(err, tc.errExpected) {
				t.Errorf("expected error %v, got %v", tc.errExpected, err)
			}
		})
	}
}

type SentTestCase struct {
	name        string
	SenderID    int
	ReceiverID  int
	Amount      float64
	errExpected error
}

type GetByReceiverIDTestCase struct {
	name       string
	ReceiverID int
	stat       int
}

type GetBySenderIDTestCase struct {
	name     string
	SenderID int
	stat     int
}

type UpdateTestCase struct {
	name        string
	ReceiverID  int
	GiftCartID  int
	status      string
	errExpected error
}
