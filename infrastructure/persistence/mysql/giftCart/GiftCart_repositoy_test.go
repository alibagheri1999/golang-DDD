package giftCart_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"remote-task/domain/giftCart/entity"
	"remote-task/infrastructure/persistence/mysql/giftCart"
	"testing"
	"time"
)

func Test_Gift_Cart_Repo(t *testing.T) {
	err := godotenv.Load("/Users/alibagheri/GolandProjects/remote-task/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbDriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	log.Println(dbDriver, host, password, user, dbname, port)
	const dbdriver = "mysql"
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?collation=utf8mb4_unicode_ci&parseTime=True", user, password, host, port, dbname)
	log.Println(DBURL)
	conn, err := sql.Open(dbdriver, DBURL)
	if err != nil {
		t.Fatal(err)
	}
	if pingErr := conn.Ping(); pingErr != nil {
		t.Fatal(pingErr)
	} else {
		log.Println("connection is ok")
	}
	log.Println(conn)
	repo := giftCart.NewGiftCardRepository(conn)
	giftCart := &entity.GiftCard{
		SenderID:   1,
		ReceiverID: 2,
		Amount:     2,
		Status:     "sent",
		CreatedAt:  time.Now(),
	}
	err = repo.CreateGiftCard(giftCart)
	if err != nil {
		t.Fatal(err)
	}
	gcs1, err1 := repo.GetGiftCardsByReceiverID(2)
	if err1 != nil {
		t.Fatal(err1)
	}
	log.Println(gcs1)
	gcs2, err2 := repo.GetGiftCardsBySenderID(1)
	if err2 != nil {
		t.Fatal(err2)
	}
	log.Println(gcs2)
	g, err3 := repo.GetGiftCardByID(1)
	if err3 != nil {
		t.Fatal(err3)
	}
	log.Println(g)
	gcs3, err4 := repo.GetGiftCardsByStatus("sent")
	if err4 != nil {
		t.Fatal(err4)
	}
	log.Println(gcs3)
	err4 = repo.UpdateGiftCardStatus(2, "accept")
	if err4 != nil {
		t.Fatal(err4)
	}
}
