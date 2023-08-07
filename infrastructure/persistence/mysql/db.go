package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	giftRepo "remote-task/domain/giftCart/repository"
	userRepo "remote-task/domain/user/repository"
	"remote-task/infrastructure/persistence/mysql/giftCart"
	"remote-task/infrastructure/persistence/mysql/user"
)

type Repositories struct {
	GiftCart giftRepo.GiftCardRepository
	User     userRepo.UserRepository
	db       *sql.DB
}

func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%v)/%s?collation=utf8mb4_unicode_ci&parseTime=True",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName,
	)
	const Dialect = "mysql"
	db, err := sql.Open(Dialect, dsn)
	if err != nil {
		fmt.Println("db connection failed", err)
	}
	if pingErr := db.Ping(); pingErr != nil {
		fmt.Println("Ping", pingErr)
	} else {
		fmt.Println("connection is ok")
	}
	return &Repositories{
		GiftCart: giftCart.NewGiftCardRepository(db),
		User:     user.NewUserRepository(db),
		db:       db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

func (s *Repositories) Ping() error {
	if err := s.db.Ping(); err != nil {
		return err
	}

	return nil
}
