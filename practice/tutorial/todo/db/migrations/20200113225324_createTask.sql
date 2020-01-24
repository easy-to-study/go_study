-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS task (
    id INT UNSIGNED NOT NULL,
    created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO task (
    id, 
    created_at,
    updated_at,
    title
) values (
    '1',
    '2019-05-02',
    '2019-05-02',
    'test'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE task;
