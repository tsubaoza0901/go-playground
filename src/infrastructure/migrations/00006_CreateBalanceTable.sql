
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS balance (
    user_id int(11) NOT NULL COMMENT 'ユーザーID',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL,
    amount int(11) NOT NULL COMMENT '金額',
    PRIMARY KEY(user_id)
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE balance;
