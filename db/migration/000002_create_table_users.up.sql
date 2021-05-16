BEGIN;

CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp without time zone NOT NULL DEFAULT NOW(),
    name text
);

COMMIT;
