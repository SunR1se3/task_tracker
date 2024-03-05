CREATE TABLE IF NOT EXISTS pool (
    id uuid NOT NULL,
    title VARCHAR(256) NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS project_pool (
    project_id uuid NOT NULL,
    pool_id uuid NOT NULL
);

