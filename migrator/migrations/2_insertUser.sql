-- +migrate Up
INSERT INTO users (username, email, created_at) VALUES
                                        ('user1', 'user1@example.com', CURRENT_TIMESTAMP),
                                        ('user2', 'user2@example.com', CURRENT_TIMESTAMP),
                                        ('user3', 'user3@example.com', CURRENT_TIMESTAMP),
                                        ('user4', 'user4@example.com', CURRENT_TIMESTAMP);