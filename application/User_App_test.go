package application_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"remote-task/application"
	"remote-task/domain/user/entity"
	"testing"
	"time"
)

var (
	GetUserByIDRepo func(c context.Context, id int) (*entity.User, error)
)

type fakeUserRepo struct{}

func (fg *fakeUserRepo) GetByID(c context.Context, id int) (*entity.User, error) {
	return GetUserByIDRepo(c, id)
}

var userAppFake application.UserAppInterface = &fakeUserRepo{}

func TestGetUserByID_Success(t *testing.T) {
	c := context.Background()

	GetUserByIDRepo = func(c context.Context, id int) (*entity.User, error) {
		return &entity.User{
			ID:        1,
			Username:  "ali",
			Email:     "ali@bagheri.com",
			CreatedAt: time.Now(),
		}, nil
	}
	r, err := userAppFake.GetByID(c, 1)
	assert.Nil(t, err)
	assert.EqualValues(t, r.Username, "ali")
	assert.EqualValues(t, r.Email, "ali@bagheri.com")
}
