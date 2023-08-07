CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert four default users
INSERT INTO users (username, email) VALUES
                                        ('user1', 'user1@example.com'),
                                        ('user2', 'user2@example.com'),
                                        ('user3', 'user3@example.com'),
                                        ('user4', 'user4@example.com');