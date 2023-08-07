package user

import (
	"database/sql"
	"remote-task/domain/giftCart/entity"
	"remote-task/domain/user/repository"
	"sync"
)

// giftCardRepositoryImpl is an implementation of the GiftCardRepository
type userRepositoryImpl struct {
	db *sql.DB
	mu sync.Mutex
}

// NewUserRepository creates a new instance of giftCardRepositoryImpl
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// CreateGiftCard creates a new gift card in the database
func (r *userRepositoryImpl) CreateGiftCard(giftCard *entity.GiftCard) error {
	return nil
}
