
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users (id, first_name, last_name, age, grade_id, email_address)
VALUES  (1, '太郎', '山田', 25, 6, "xxxxxx@google.com"),
        (2, '美子', '向井', 50, 5, "tttt@google.com");


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM users;
