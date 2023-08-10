package repository_test

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"remote-task/domain/user/repository"
	"remote-task/domain/user/userConst"
	"remote-task/infrastructure/persistence/mysql"
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
	repo, err1 := mysql.NewRepositories(user, password, port, host, dbname)
	if err1 != nil {
		t.Fatal(err1)
	}
	UserRepo := repository.NewUserRepository(repo)
	getByIDTestCase := []GetByIDTestCase{
		{
			name:        "user exist",
			id:          1,
			errExpected: nil,
		},
		{
			name:        "user not exist",
			id:          0,
			errExpected: userConst.ERR_NOT_FOUND,
		},
	}
	for _, tc := range getByIDTestCase {
		t.Run(tc.name, func(t *testing.T) {
			_, err := UserRepo.GetUserByID(context.Background(), tc.id)
			if !errors.Is(err, tc.errExpected) {
				t.Errorf("expected error %v, got %v", tc.errExpected, err)
			}
		})
	}
}

type GetByIDTestCase struct {
	name        string
	id          int
	errExpected error
}
