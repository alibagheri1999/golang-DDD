package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"math"
	"os"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/repository"
	"remote-task/domain/user/userConst"
	"remote-task/infrastructure/persistence/mysql"
	"remote-task/utilities"
	"testing"
)

func Test_Gift_Cart_Repo(t *testing.T) {
	err := godotenv.Load("/Users/alibagheri/GolandProjects/remote-task/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	const dbdriver = "mysql"
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?collation=utf8mb4_unicode_ci&parseTime=True", user, password, host, port, dbname)
	log.Println(DBURL)
	conn, err := sql.Open(dbdriver, DBURL)
	if err != nil {
		t.Fatal(err)
	}
	conn.SetMaxOpenConns(0)
	conn.SetMaxIdleConns(2)

	if pingErr := conn.Ping(); pingErr != nil {
		log.Println("Err mysql ping", pingErr)
	} else {
		log.Println("Success mysql connection is ok")
	}

	var skip string
	var maxConnections int
	maxConErr := conn.QueryRow(utilities.SHOW_VARS_CONNECTION).Scan(&skip, &maxConnections)
	if maxConErr != nil {
		log.Println("Err mysql getting the max_connections", maxConErr)
	}
	maxConnections = int(math.Floor(float64(maxConnections) * 0.9))
	if maxConnections == 0 {
		maxConnections = 100
	}
	conn.SetMaxOpenConns(maxConnections)

	var waitTimeout int
	waitErr := conn.QueryRow(utilities.SHOW_VARS_TIMEOUT).Scan(&skip, &waitTimeout)
	if waitErr != nil {
		log.Println("Err mysql getting the wait_timeout", waitErr)
	}
	if waitTimeout == 0 {
		waitTimeout = 180
	}
	waitTimeout = int(math.Min(float64(waitTimeout), 180))

	repo := &mysql.Repositories{
		Db:         conn,
		Statements: make(map[string]*sql.Stmt),
	}
	GiftCarRepo := repository.NewGiftCardRepository(repo)
	sentGiftTestCase := []SentGiftTestCase{
		{
			name:        "ReceiverID out of range",
			ReceiverID:  0,
			SenderID:    1,
			Amount:      2,
			errExpected: userConst.ERR_NOT_FOUND,
		},
		{
			name:        "SenderID out of range",
			ReceiverID:  1,
			SenderID:    0,
			Amount:      1,
			errExpected: userConst.ERR_NOT_FOUND,
		},
		{
			name:        "created successfully",
			ReceiverID:  1,
			SenderID:    1,
			Amount:      1,
			errExpected: nil,
		},
	}
	for _, tc := range sentGiftTestCase {
		t.Run(tc.name, func(t *testing.T) {
			giftCart := &DTO.SendGiftCartRequest{
				SenderID:   tc.SenderID,
				ReceiverID: tc.ReceiverID,
				Amount:     tc.Amount,
			}
			err := GiftCarRepo.CreateGiftCard(context.Background(), giftCart)
			if !errors.Is(err, tc.errExpected) {
				t.Errorf("expected error %v, got %v", tc.errExpected, err)
			}
		})
	}
	//gcs1, err1 := GiftCarRepo.GetGiftCardsByReceiverID(context.Background(), 2, 'sent')
	//if err1 != nil {
	//	t.Fatal(err1)
	//}
	//log.Println(gcs1)
	//gcs2, err2 := GiftCarRepo.GetGiftCardsBySenderID(context.Background(), 1)
	//if err2 != nil {
	//	t.Fatal(err2)
	//}
	//log.Println(gcs2)
	//g, err3 := GiftCarRepo.GetGiftCardByID(context.Background(), 1)
	//if err3 != nil {
	//	t.Fatal(err3)
	//}
	//log.Println(g)
	//gcs3, err4 := GiftCarRepo.GetGiftCardsByStatus(context.Background(), "sent")
	//if err4 != nil {
	//	t.Fatal(err4)
	//}
	//log.Println(gcs3)
	//err4 = GiftCarRepo.UpdateGiftCardStatus(context.Background(), 2, "accept")
	//if err4 != nil {
	//	t.Fatal(err4)
	//}
}

type SentGiftTestCase struct {
	name        string
	SenderID    int
	ReceiverID  int
	Amount      float64
	errExpected error
}
