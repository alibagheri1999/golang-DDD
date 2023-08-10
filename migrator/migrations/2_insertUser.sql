-- +migrate Up
INSERT INTO users (username, email) VALUES
                                        ('user1', 'user1@example.com'),
                                        ('user2', 'user2@example.com'),
                                        ('user3', 'user3@example.com'),
                                        ('user4', 'user4@example.com');