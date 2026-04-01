CREATE TABLE IF NOT EXISTS followers (
    user_id bigint NOT NULL,
    follower_id bigint NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
    updated_at timestamp(0) with time zone DEFAULT NULL,
    deleted_at timestamp(0) with time zone DEFAULT NULL,
    PRIMARY KEY (user_id, follower_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (follower_id) REFERENCES users(id)
);
