CREATE TABLE IF NOT EXISTS task (
    id uuid NOT NULL,
    title VARCHAR(256) NOT NULL,
    description text NOT NULL,
    tag_id uuid DEFAULT NULL,
    tags jsonb DEFAULT NULL,
    difficulty int NOT NULL,
    executor_id uuid DEFAULT NULL,
    executor jsonb DEFAULT NULL,
    author_id uuid NOT NULL,
    author jsonb NOT NULL,
    status int,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    PRIMARY KEY(id)
);