CREATE TABLE users
(
    id           INT          AUTO_INCREMENT PRIMARY KEY,
    uid          VARCHAR(256) NOT NULL,
    name         VARCHAR(128) NOT NULL,
    thumbnail    VARCHAR(256),
    created_at   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_users ON users(uid);