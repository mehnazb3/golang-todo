BEGIN;

CREATE TABLE IF NOT EXISTS todos(
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp without time zone NOT NULL DEFAULT NOW(),
    text text,
    done boolean,
    user_id uuid
);

COMMIT;