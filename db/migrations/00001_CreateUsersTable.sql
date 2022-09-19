
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS users (
    id int(11) NOT NULL AUTO_INCREMENT,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL,
    first_name varchar(128) NOT NULL COMMENT "名",
    last_name varchar(128) NOT NULL COMMENT "姓",
    age int(11) NOT NULL COMMENT "年齢",
    grade_id int(11) NOT NULL COMMENT 'グレードID',
    email_address varchar(128) NOT NULL COMMENT "メールアドレス",
    PRIMARY KEY(id)
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
