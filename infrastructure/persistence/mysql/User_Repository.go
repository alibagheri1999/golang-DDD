package mysql

import (
	"context"
	"remote-task/domain/user/entity"
	"remote-task/domain/user/repository"
	"remote-task/domain/user/userConst"
	"remote-task/utilities"
	"sync"
)

// giftCardRepositoryImpl is an implementation of the GiftCardRepository
type UserRepositoryImpl struct {
	mysqlRepo *Repositories
	mu        sync.Mutex
}

// NewUserRepository creates a new instance of giftCardRepositoryImpl
func NewUserRepository(mysqlRepo *Repositories) repository.UserRepository {
	return &UserRepositoryImpl{
		mysqlRepo: mysqlRepo,
	}
}

// GetUserByID get a user with its id
func (r *UserRepositoryImpl) GetUserByID(c context.Context, id int) (*entity.User, error) {
	query := utilities.GET_USER_BY_ID
	stmt := r.mysqlRepo.stmt("stmtGetUserByID")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			return nil, err
		}
		r.mysqlRepo.setStmt("stmtGetUserByID", ps)
		stmt = ps
	}
	user := &entity.User{}
	err2 := stmt.QueryRowContext(c, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err2 != nil {
		return nil, userConst.ERR_NOT_FOUND
	}
	return user, nil
}
