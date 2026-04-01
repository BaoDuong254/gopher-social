CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    content text NOT NULL,
    title text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
    updated_at timestamp(0) with time zone DEFAULT NULL,
    deleted_at timestamp(0) with time zone DEFAULT NULL
)
