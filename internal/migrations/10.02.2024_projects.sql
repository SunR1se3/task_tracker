CREATE TABLE IF NOT EXISTS projects (
    id uuid NOT NULL,
    title VARCHAR(500) NOT NULL,
    description VARCHAR(500) NOT NULL,
    consumer VARCHAR(500) DEFAULT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS user_project (
    user_id uuid NOT NULL,
    project_id uuid NOT NULL,
    project_role_id uuid DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS project_roles (
    id uuid NOT NULL,
    title VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
)