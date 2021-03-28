CREATE TABLE IF NOT EXISTS event
(
    sequence BIGSERIAL NOT NULL,
    id UUID NOT NULL,
    aggregate TEXT NOT NULL,
    type TEXT NOT NULL,
    data JSONB,
    meta JSON,
    UNIQUE (id, sequence),
)