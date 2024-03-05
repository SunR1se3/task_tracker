CREATE TABLE IF NOT EXISTS sprints (
    id uuid NOT NULL,
    title VARCHAR(256) NOT NULL,
    start_date TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    end_date TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS sprint_task (
    sprint_id uuid NOT NULL,
    task_id uuid NOT NULL
);