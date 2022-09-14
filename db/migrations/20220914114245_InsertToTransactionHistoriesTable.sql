
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO deal_histories (id, user_id, item_name, amount)
VALUES  (1, 1, "電車代", 1000);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM deal_histories;

