-- Drop ENUM type for priority if it exists
DROP TYPE IF EXISTS PriorityEnum;
-- Create ENUM type for priority
CREATE TYPE PriorityEnum AS ENUM ('LOW', 'MEDIUM', 'HIGH', 'URGENT');

-- Drop ENUM type for status if it exists
DROP TYPE IF EXISTS StatusEnum;
-- Create ENUM type for status
CREATE TYPE StatusEnum AS ENUM ('BACKLOG', 'TODO', 'IN_PROGRESS', 'DONE');

-- Drop the task table if it exists
DROP TABLE IF EXISTS task;
-- Create the task table
CREATE TABLE task (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    priority PriorityEnum NOT NULL,
    assignee BIGINT,
    status StatusEnum NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

-- Drop the team table if it exists
DROP TABLE IF EXISTS team;
-- Create the team table
CREATE TABLE team (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

-- Drop the user table if it exists
DROP TABLE IF EXISTS "user";
-- Create the user table
CREATE TABLE "user" (
    id BIGINT PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    team_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

-- Drop the event table if it exists
DROP TABLE IF EXISTS event;
-- Create the event table
CREATE TABLE event (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    time TIMESTAMP NOT NULL,
    duration INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);