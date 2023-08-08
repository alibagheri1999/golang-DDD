package mysql

import (
	"context"
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
	mysqlRepo *Repositories
	mu        sync.Mutex
}

// NewGiftCardRepository creates a new instance of giftCardRepositoryImpl
func NewGiftCardRepository(repo *Repositories) repository.GiftCardRepository {
	return &giftCardRepositoryImpl{
		mysqlRepo: repo,
	}
}

// CreateGiftCard creates a new gift card in the database
func (r *giftCardRepositoryImpl) CreateGiftCard(c context.Context, giftCard *entity.GiftCard) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.INSERT_GIFT
	stmt := r.mysqlRepo.stmt("stmtCreateGiftCart")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			return err
		}
		r.mysqlRepo.setStmt("stmtCreateGiftCart", ps)
		stmt = ps
	}
	_, err := stmt.ExecContext(c,
		giftCard.SenderID, giftCard.ReceiverID, giftCard.Amount, giftCard.Status, giftCard.CreatedAt,
	)
	if err != nil {
		log.Println(err)
		return giftCartConst.ERR_CREATE_GIFT_CART
	}
	return nil
}

// GetGiftCardByID retrieves a gift card from the database by its ID
func (r *giftCardRepositoryImpl) GetGiftCardByID(c context.Context, id int) (*entity.GiftCard, error) {
	query := utilities.GET_BY_ID
	stmt := r.mysqlRepo.stmt("stmtGetGiftCardByID")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			log.Printf("prepare context error: %v\n", err)
			return nil, err

		}
		r.mysqlRepo.setStmt("stmtGetGiftCardByID", ps)
		stmt = ps
	}
	giftCard := &entity.GiftCard{}
	err2 := stmt.QueryRowContext(c, id).Scan(&giftCard.ID, &giftCard.SenderID, &giftCard.ReceiverID, &giftCard.Amount, &giftCard.Status, &giftCard.CreatedAt)
	if err2 != nil {
		log.Printf("prepare context error: %v\n", err2)
		return nil, giftCartConst.ERR_NOT_FOUND
	}
	return giftCard, nil
}

// GetGiftCardsByReceiverID retrieves a gift card from the database by its ReceiverID
func (r *giftCardRepositoryImpl) GetGiftCardsByReceiverID(c context.Context, receiverID int) ([]entity.GiftCardJoinUserByReceiver, error) {
	query := utilities.GET_GIFT_BY_RECEIVER_ID
	stmt := r.mysqlRepo.stmt("stmtGetGiftCardsByReceiverID")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			log.Printf("prepare context error: %v\n", err)
			return nil, err

		}
		r.mysqlRepo.setStmt("stmtGetGiftCardsByReceiverID", ps)
		stmt = ps
	}
	rows, err := stmt.QueryContext(c, receiverID)
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
func (r *giftCardRepositoryImpl) GetGiftCardsBySenderID(c context.Context, senderID int) ([]entity.GiftCardJoinUserBySender, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.GET_GIFT_BY_SENDER_ID
	stmt := r.mysqlRepo.stmt("stmtGetGiftCardsBySenderID")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			log.Printf("prepare context error: %v\n", err)
			return nil, err

		}
		r.mysqlRepo.setStmt("stmtGetGiftCardsBySenderID", ps)
		stmt = ps
	}
	rows, err := stmt.QueryContext(c, senderID)
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
func (r *giftCardRepositoryImpl) GetGiftCardsByStatus(c context.Context, status string) ([]entity.GiftCard, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.GET_BY_STATUS
	stmt := r.mysqlRepo.stmt("stmtGetGiftCardsByStatus")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			log.Printf("prepare context error: %v\n", err)
			return nil, err

		}
		r.mysqlRepo.setStmt("stmtGetGiftCardsByStatus", ps)
		stmt = ps
	}
	rows, err := stmt.QueryContext(c, status)
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
func (r *giftCardRepositoryImpl) UpdateGiftCardStatus(c context.Context, id int, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.UPDATE_STATUS
	stmt := r.mysqlRepo.stmt("stmtUpdateGiftCart")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			return err
		}
		r.mysqlRepo.setStmt("stmtUpdateGiftCart", ps)
		stmt = ps
	}
	res, err := stmt.ExecContext(c,
		status, id)
	if err != nil {
		log.Println(err)
		return giftCartConst.ERR_UPDATE_GIFT_CART
	}
	_, err1 := res.RowsAffected()
	if err1 != nil {
		return giftCartConst.ERR_UPDATE_GIFT_CART
	}
	return nil
}
