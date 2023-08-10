-- +migrate Up
INSERT INTO gift_cards (sender_id, receiver_id, amount, status)
VALUES
    (1, 2, 50.00, 'accept'),
    (1, 2, 50.00, 'reject'),
    (1, 2, 50.00, 'sent'),
    (1, 3, 75.00, 'accept'),
    (1, 3, 75.00, 'reject'),
    (1, 3, 75.00, 'sent'),
    (1, 4, 30.00, 'accept');
    (1, 4, 30.00, 'reject');
    (1, 4, 30.00, 'sent');
    (2, 4, 30.00, 'reject');
    (2, 4, 30.00, 'sent');
    (2, 4, 30.00, 'accept');
