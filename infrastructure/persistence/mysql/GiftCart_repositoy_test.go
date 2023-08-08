package mysql_test

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"math"
	"os"
	"remote-task/domain/giftCart/entity"
	"remote-task/infrastructure/persistence/mysql"
	"remote-task/utilities"
	"testing"
	"time"
)

func Test_Gift_Cart_Repo(t *testing.T) {
	err := godotenv.Load("/Users/alibagheri/GolandProjects/remote-task/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
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
	GiftCarRepo := mysql.NewGiftCardRepository(repo)
	giftCart := &entity.GiftCard{
		SenderID:   1,
		ReceiverID: 2,
		Amount:     2,
		Status:     "sent",
		CreatedAt:  time.Now(),
	}
	err = GiftCarRepo.CreateGiftCard(context.Background(), giftCart)
	if err != nil {
		t.Fatal(err)
	}
	gcs1, err1 := GiftCarRepo.GetGiftCardsByReceiverID(context.Background(), 2)
	if err1 != nil {
		t.Fatal(err1)
	}
	log.Println(gcs1)
	gcs2, err2 := GiftCarRepo.GetGiftCardsBySenderID(context.Background(), 1)
	if err2 != nil {
		t.Fatal(err2)
	}
	log.Println(gcs2)
	g, err3 := GiftCarRepo.GetGiftCardByID(context.Background(), 1)
	if err3 != nil {
		t.Fatal(err3)
	}
	log.Println(g)
	gcs3, err4 := GiftCarRepo.GetGiftCardsByStatus(context.Background(), "sent")
	if err4 != nil {
		t.Fatal(err4)
	}
	log.Println(gcs3)
	err4 = GiftCarRepo.UpdateGiftCardStatus(context.Background(), 2, "accept")
	if err4 != nil {
		t.Fatal(err4)
	}
}
