-- +migrate Up
CREATE TABLE gift_cards (
                           id INT AUTO_INCREMENT PRIMARY KEY,
                           sender_id INT NOT NULL,
                           receiver_id INT NOT NULL,
                           amount DECIMAL(10, 2) NOT NULL,
                           status VARCHAR(50) NOT NULL,
                           created_at TIMESTAMP NOT NULL ,
                           FOREIGN KEY (sender_id) REFERENCES users (id),
                           FOREIGN KEY (receiver_id) REFERENCES users (id)
);