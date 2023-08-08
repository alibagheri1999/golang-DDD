package application

import (
	"context"
	"remote-task/domain/user/entity"
	"remote-task/domain/user/repository"
)

type userApp struct {
	ur repository.UserRepository
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	GetUserByID(c context.Context, id int) (*entity.User, error)
}

func (u *userApp) GetUserByID(c context.Context, id int) (*entity.User, error) {
	return u.ur.GetUserByID(c, id)
}
