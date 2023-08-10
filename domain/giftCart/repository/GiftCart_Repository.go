package repository

import (
	"context"
	"database/sql"
	"log"
	"remote-task/domain/giftCart/DTO"
	"remote-task/domain/giftCart/entity"
	"remote-task/domain/giftCart/giftCartConst"
	"remote-task/domain/giftCart/param"
	"remote-task/infrastructure/persistence/mysql"
	"remote-task/utilities"
	"sync"
	"time"
)

// GiftCardRepositoryImpl is an implementation of the GiftCardRepository
type GiftCardRepositoryImpl struct {
	mysqlRepo *mysql.Repositories
	mu        sync.Mutex
}

// NewGiftCardRepository creates a new instance of GiftCardRepositoryImpl
func NewGiftCardRepository(repo *mysql.Repositories) GiftCardRepository {
	return &GiftCardRepositoryImpl{
		mysqlRepo: repo,
	}
}

// Create creates a new gift card in the database
func (r *GiftCardRepositoryImpl) Create(c context.Context, giftCard *DTO.SendGiftCartRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	gc := entity.GiftCard{
		CreatedAt:  time.Now(),
		Status:     "sent",
		Amount:     giftCard.Amount,
		SenderID:   giftCard.SenderID,
		ReceiverID: giftCard.ReceiverID,
	}
	query := utilities.INSERT_GIFT
	stmt := r.mysqlRepo.Stmt("stmtCreateGiftCart")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			return err
		}
		r.mysqlRepo.SetStmt("stmtCreateGiftCart", ps)
		stmt = ps
	}
	_, err := stmt.ExecContext(c,
		&gc.SenderID, &gc.ReceiverID, &gc.Amount, &gc.Status, &gc.CreatedAt,
	)
	if err != nil {
		log.Println(err)
		return giftCartConst.ERR_CREATE_GIFT_CART
	}
	return nil
}

// GetByID retrieves a gift card from the database by its ID
func (r *GiftCardRepositoryImpl) GetByID(c context.Context, id int) (*entity.GiftCard, error) {
	query := utilities.GET_GIDT_BY_ID
	stmt := r.mysqlRepo.Stmt("stmtGetGiftCardByID")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			log.Printf("prepare context error: %v\n", err)
			return nil, err

		}
		r.mysqlRepo.SetStmt("stmtGetGiftCardByID", ps)
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

// GetByReceiverID retrieves a gift card from the database by its ReceiverID
func (r *GiftCardRepositoryImpl) GetByReceiverID(c context.Context, receiverID int, stat int) ([]param.GiftCardJoinUserByReceiver, error) {
	query := utilities.GET_GIFT_BY_RECEIVER_ID
	var status string
	stmt := r.mysqlRepo.Stmt("stmtGetGiftCardsByReceiverID")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			log.Printf("prepare context error: %v\n", err)
			return nil, err

		}
		r.mysqlRepo.SetStmt("stmtGetGiftCardsByReceiverID", ps)
		stmt = ps
	}
	switch stat {
	case 1:
		status = "accept"
	case 2:
		status = "reject"
	case 3:
		status = "sent"
	}
	rows, err := stmt.QueryContext(c, receiverID, status)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}(rows)
	var giftCards []param.GiftCardJoinUserByReceiver
	for rows.Next() {
		giftCard := &param.GiftCardJoinUserByReceiver{}
		err := rows.Scan(&giftCard.GiftCardID, &giftCard.SenderID, &giftCard.UserID, &giftCard.ReceiverName, &giftCard.ReceiverEmail,
			&giftCard.ReceiverID, &giftCard.Amount, &giftCard.Status, &giftCard.GiftCardCreatedAt)
		if err != nil {
			return nil, err
		}
		giftCards = append(giftCards, *giftCard)
	}

	return giftCards, nil
}

// GetBySenderID retrieves a gift card from the database by its SenderID
func (r *GiftCardRepositoryImpl) GetBySenderID(c context.Context, senderID int, stat int) ([]param.GiftCardJoinUserBySender, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := utilities.GET_GIFT_BY_SENDER_ID
	switch stat {
	case 1:
		query += utilities.ADD_ACCEPT_STATUS
	case 2:
		query += utilities.ADD_REJECT_STATUS
	case 3:
		query += utilities.ADD_SENT_STATUS
	}
	stmt := r.mysqlRepo.Stmt("stmtGetGiftCardsBySenderID")
	ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
	if err != nil {
		log.Printf("prepare context error: %v\n", err)
		return nil, err

	}
	r.mysqlRepo.SetStmt("stmtGetGiftCardsBySenderID", ps)
	stmt = ps
	rows, err := stmt.QueryContext(c, senderID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}(rows)
	var giftCards []param.GiftCardJoinUserBySender
	for rows.Next() {
		giftCard := &param.GiftCardJoinUserBySender{}
		err := rows.Scan(&giftCard.GiftCardID, &giftCard.SenderID, &giftCard.UserID, &giftCard.SenderName, &giftCard.SenderEmail,
			&giftCard.ReceiverID, &giftCard.Amount, &giftCard.Status, &giftCard.GiftCardCreatedAt)
		if err != nil {
			return nil, giftCartConst.ERR_NOT_FOUND
		}
		giftCards = append(giftCards, *giftCard)
	}

	return giftCards, nil
}

// UpdateStatus update a gift card in the database in terms of status
func (r *GiftCardRepositoryImpl) UpdateStatus(c context.Context, receiverID int, giftCartID int, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, err := r.GetByID(c, giftCartID)
	if err != nil {
		return giftCartConst.ERR_NOT_FOUND
	}
	query := utilities.UPDATE_STATUS
	stmt := r.mysqlRepo.Stmt("stmtUpdateGiftCart")
	if stmt == nil {
		ps, err := r.mysqlRepo.Db.PrepareContext(c, query)
		if err != nil {
			return err
		}
		r.mysqlRepo.SetStmt("stmtUpdateGiftCart", ps)
		stmt = ps
	}
	if status != "accept" && status != "reject" {
		return giftCartConst.ERR_UPDATE_GIFT_CART
	}
	res, err1 := stmt.ExecContext(c,
		status, receiverID, giftCartID)
	if err1 != nil {
		log.Println(err1)
		return giftCartConst.ERR_UPDATE_GIFT_CART
	}
	_, err2 := res.RowsAffected()
	if err2 != nil {
		return giftCartConst.ERR_UPDATE_GIFT_CART
	}
	return nil
}
