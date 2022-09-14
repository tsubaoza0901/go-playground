
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO grades (id, name)
VALUES  (1, '1級'),
        (2, '2級'),
        (3, '3級'),
        (4, '4級'),
        (5, '5級'),
        (6, '等級なし');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM grades;