package repository

import (
	"context"
	"remote-task/domain/user/entity"
)

type UserRepository interface {
	GetByID(c context.Context, id int) (*entity.User, error)
}
