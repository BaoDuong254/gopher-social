CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email citext NOT NULL UNIQUE,
    username varchar(255) NOT NULL UNIQUE,
    password bytea NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
    updated_at timestamp(0) with time zone DEFAULT NULL,
    deleted_at timestamp(0) with time zone DEFAULT NULL
);
