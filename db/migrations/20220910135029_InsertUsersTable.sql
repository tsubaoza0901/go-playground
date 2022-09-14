
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users (id, first_name, last_name, age, grade_id)
VALUES  (1, '太郎', '山田', 25, 6);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM users;
