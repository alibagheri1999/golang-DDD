package utilities

const (
	INSERT_GIFT           = "INSERT INTO gift_cards (sender_id, receiver_id, amount, status, created_at) VALUES (?, ?, ?, ?, ?)"
	GET_BY_ID             = "SELECT id, sender_id, receiver_id, amount, status, created_at FROM gift_cards WHERE id = ? LIMIT 1"
	GET_BY_STATUS         = "SELECT id, sender_id, receiver_id, amount, status, created_at FROM gift_cards WHERE status = ?"
	UPDATE_STATUS         = "UPDATE gift_cards SET status = ? WHERE id = ?"
	GET_GIFT_BY_SENDER_ID = `
		SELECT gc.id, gc.sender_id, u.id, u.username, u.email,
		gc.receiver_id, gc.amount, gc.status, gc.created_at
		FROM gift_cards AS gc
		JOIN users AS u ON gc.sender_id = u.id
		WHERE gc.sender_id = ?
	`
	GET_GIFT_BY_RECEIVER_ID = `
		SELECT gc.id, gc.sender_id, u.id, u.username, u.email,
		gc.receiver_id, gc.amount, gc.status, gc.created_at
		FROM gift_cards AS gc
		JOIN users AS u ON gc.receiver_id = u.id
		WHERE gc.receiver_id = ?
	`
)
