package giftCart

import (
	"database/sql"
	"log"
	"remote-task/domain/giftCart/entity"
	"remote-task/domain/giftCart/giftCartConst"
	"remote-task/domain/giftCart/repository"
	"remote-task/utilities"
	"sync"
)

// giftCardRepositoryImpl is an implementation of the GiftCardRepository
type giftCardRepositoryImpl struct {
	db *sql.DB
	mu sync.Mutex
}

// NewGiftCardRepository creates a new instance of giftCardRepositoryImpl
func NewGiftCardRepository(db *sql.DB) repository.GiftCardRepository {
	return &giftCardRepositoryImpl{
		db: db,
	}
}

// CreateGiftCard creates a new gift card in the database
func (r *giftCardRepositoryImpl) CreateGiftCard(giftCard *entity.GiftCard) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.INSERT_GIFT
	_, err := r.db.Exec(query, giftCard.SenderID, giftCard.ReceiverID, giftCard.Amount, giftCard.Status, giftCard.CreatedAt)
	if err != nil {
		log.Println(err)
		return giftCartConst.ERR_CREATE_GIFT_CART
	}
	return nil
}

// GetGiftCardByID retrieves a gift card from the database by its ID
func (r *giftCardRepositoryImpl) GetGiftCardByID(id int) (*entity.GiftCard, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.GET_BY_ID
	row := r.db.QueryRow(query, id)

	giftCard := &entity.GiftCard{}
	err := row.Scan(&giftCard.ID, &giftCard.SenderID, &giftCard.ReceiverID, &giftCard.Amount, &giftCard.Status, &giftCard.CreatedAt)
	if err != nil {
		return nil, giftCartConst.ERR_NOT_FOUND
	}
	return giftCard, nil
}

// GetGiftCardsByReceiverID retrieves a gift card from the database by its ReceiverID
func (r *giftCardRepositoryImpl) GetGiftCardsByReceiverID(receiverID int) ([]entity.GiftCardJoinUserByReceiver, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.GET_GIFT_BY_RECEIVER_ID
	rows, err := r.db.Query(query, receiverID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}(rows)
	var giftCards []entity.GiftCardJoinUserByReceiver
	for rows.Next() {
		log.Println(rows)
		giftCard := &entity.GiftCardJoinUserByReceiver{}
		err := rows.Scan(&giftCard.GiftCardID, &giftCard.SenderID, &giftCard.UserID, &giftCard.ReceiverName, &giftCard.ReceiverEmail,
			&giftCard.ReceiverID, &giftCard.Amount, &giftCard.Status, &giftCard.GiftCardCreatedAt)
		if err != nil {
			return nil, err
		}
		giftCards = append(giftCards, *giftCard)
	}

	return giftCards, nil
}

// GetGiftCardsBySenderID retrieves a gift card from the database by its SenderID
func (r *giftCardRepositoryImpl) GetGiftCardsBySenderID(senderID int) ([]entity.GiftCardJoinUserBySender, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.GET_GIFT_BY_SENDER_ID
	rows, err := r.db.Query(query, senderID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}(rows)
	var giftCards []entity.GiftCardJoinUserBySender
	for rows.Next() {
		giftCard := &entity.GiftCardJoinUserBySender{}
		err := rows.Scan(&giftCard.GiftCardID, &giftCard.SenderID, &giftCard.UserID, &giftCard.SenderName, &giftCard.SenderEmail,
			&giftCard.ReceiverID, &giftCard.Amount, &giftCard.Status, &giftCard.GiftCardCreatedAt)
		if err != nil {
			return nil, giftCartConst.ERR_NOT_FOUND
		}
		giftCards = append(giftCards, *giftCard)
	}

	return giftCards, nil
}

// GetGiftCardsByStatus retrieves all gift cards from the database by their Status
func (r *giftCardRepositoryImpl) GetGiftCardsByStatus(status string) ([]entity.GiftCard, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.GET_BY_STATUS
	rows, err := r.db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}(rows)
	var giftCards []entity.GiftCard
	for rows.Next() {
		giftCard := &entity.GiftCard{}
		err := rows.Scan(&giftCard.ID, &giftCard.SenderID, &giftCard.ReceiverID, &giftCard.Amount, &giftCard.Status, &giftCard.CreatedAt)
		if err != nil {
			return nil, giftCartConst.ERR_NOT_FOUND
		}
		giftCards = append(giftCards, *giftCard)
	}

	return giftCards, nil
}

// UpdateGiftCardStatus update a gift card in the database in terms of status
func (r *giftCardRepositoryImpl) UpdateGiftCardStatus(id int, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.UPDATE_STATUS
	result, err := r.db.Exec(query, status, id)
	if err != nil {
		return giftCartConst.ERR_UPDATE_GIFT_CART
	}
	_, err = result.RowsAffected()
	if err != nil {
		return giftCartConst.ERR_UPDATE_GIFT_CART
	}
	return nil
}
