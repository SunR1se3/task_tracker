CREATE TABLE IF NOT EXISTS positions (
    id uuid NOT NULL,
    title VARCHAR(256) NOT NULL,
    code VARCHAR(256) NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS specializations (
    id uuid NOT NULL,
    title VARCHAR(256) NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS departments (
    id uuid NOT NULL,
    title VARCHAR(256) NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS department_chief (
    department_id uuid NOT NULL,
    user_id uuid NOT NULL
);
CREATE TABLE IF NOT EXISTS department_curator (
    department_id uuid NOT NULL,
    user_id uuid NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL,
    login VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    firstname VARCHAR(256) NOT NULL,
    middlename VARCHAR(256) NOT NULL,
    lastname VARCHAR(256) NOT NULL,
    is_active BOOLEAN NOT NULL,
    account_disable_time TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    system_role INTEGER NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS user_department (
    user_id uuid NOT NULL,
    department_id uuid NOT NULL
);
CREATE TABLE IF NOT EXISTS user_position (
    user_id uuid NOT NULL,
    position_id uuid NOT NULL
);
CREATE TABLE IF NOT EXISTS user_specialization (
    user_id uuid NOT NULL,
    specialization_id uuid NOT NULL
);