
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO balance (user_id, amount)
VALUES  (1, 1000),
        (2, 4000);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM balance;

