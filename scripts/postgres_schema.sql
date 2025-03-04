-- Drop ENUM type for priority if it exists
DROP TYPE IF EXISTS PriorityEnum;
-- Create ENUM type for priority
CREATE TYPE PriorityEnum AS ENUM ('LOW', 'MEDIUM', 'HIGH', 'URGENT');

-- Drop ENUM type for status if it exists
DROP TYPE IF EXISTS StatusEnum;
-- Create ENUM type for status
CREATE TYPE StatusEnum AS ENUM ('BACKLOG', 'TODO', 'IN_PROGRESS', 'DONE');

-- Drop the tasks table if it exists
DROP TABLE IF EXISTS tasks;
-- Create the tasks table with ENUM types
CREATE TABLE IF NOT EXISTS tasks (
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