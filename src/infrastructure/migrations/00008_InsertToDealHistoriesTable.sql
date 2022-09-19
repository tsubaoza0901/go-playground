
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO deal_histories (id, user_id, item_name, amount)
VALUES  (1, 1, "チャージ", 2000),
        (2, 1, "電車代", 1000),
        (3, 2, "チャージ", 8000),
        (4, 2, "電車代", 1000),
        (5, 2, "電車代", 1000),
        (6, 2, "電車代", 1000),
        (7, 2, "電車代", 1000);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM deal_histories;

