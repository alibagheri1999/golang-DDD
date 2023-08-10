package repository_test

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	appConfig "remote-task/config"
	"remote-task/domain/user/repository"
	"remote-task/domain/user/userConst"
	"remote-task/infrastructure/persistence/mysql"
	"testing"
)

// Test_User_Repo test all methods in user repository
func Test_User_Repo(t *testing.T) {
	appConfig.Init()
	dbCfg := appConfig.Get().Mysql
	repo, err := mysql.NewRepositories(dbCfg.Username, dbCfg.Password, dbCfg.Port, dbCfg.Host, dbCfg.Name)
	if err != nil {
		t.Fatal(err)
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
			_, err := UserRepo.GetByID(context.Background(), tc.id)
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
