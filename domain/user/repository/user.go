package repository

import (
	"context"
	"remote-task/domain/user/entity"
)

type UserRepository interface {
	GetUserByID(c context.Context, id int) (*entity.User, error)
}
