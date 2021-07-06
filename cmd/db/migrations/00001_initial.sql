-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto SCHEMA public;

CREATE TABLE to_dos (
    id uuid DEFAULT gen_random_uuid() NOT NULL CONSTRAINT todos_pkey PRIMARY KEY,
    title                 varchar(255)                      NOT NULL,
    description           varchar(255)                      NOT NULL,
    due_date              timestamptz                       NOT NULL,
    is_done               BOOLEAN DEFAULT FALSE             NOT NULL,
    labels                TEXT[],
    created_at            timestamptz DEFAULT now()         NOT NULL,
    updated_at            timestamptz DEFAULT now()         NOT NULL
);

;

-- +goose Down
DROP TABLE todos;

;
