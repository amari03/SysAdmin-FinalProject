CREATE TABLE IF NOT EXISTS mood_notes (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    content text NOT NULL,
    emoji text NOT NULL,
    color text NOT NULL,
    user_id bigint NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    version integer NOT NULL DEFAULT 1
);