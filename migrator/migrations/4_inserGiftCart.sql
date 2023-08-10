-- +migrate Up
INSERT INTO gift_cards (sender_id, receiver_id, amount, status, created_at) VALUES
                                                                                    (1, 2, 50.00, 'accept', CURRENT_TIMESTAMP),
                                                                                    (1, 2, 50.00, 'reject', CURRENT_TIMESTAMP),
                                                                                    (1, 2, 50.00, 'sent', CURRENT_TIMESTAMP),
                                                                                    (1, 3, 75.00, 'accept', CURRENT_TIMESTAMP),
                                                                                    (1, 3, 75.00, 'reject', CURRENT_TIMESTAMP),
                                                                                    (1, 3, 75.00, 'sent', CURRENT_TIMESTAMP),
                                                                                    (1, 4, 30.00, 'accept', CURRENT_TIMESTAMP),
                                                                                    (1, 4, 30.00, 'reject', CURRENT_TIMESTAMP),
                                                                                    (1, 4, 30.00, 'sent', CURRENT_TIMESTAMP),
                                                                                    (2, 4, 30.00, 'reject', CURRENT_TIMESTAMP),
                                                                                    (2, 4, 30.00, 'sent', CURRENT_TIMESTAMP),
                                                                                    (2, 4, 30.00, 'accept', CURRENT_TIMESTAMP);

-- +migrate Down
DELETE FROM gift_cards WHERE `sender_id`=1;
DELETE FROM gift_cards WHERE `sender_id`=2;


