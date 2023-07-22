CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMP,
    completed boolean NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

INSERT INTO tasks (title, description, due_date, completed, created_at, updated_at)
VALUES ('Finish report', 'Write a report for the project', '2023-07-25 15:00:00', false, NOW(), NOW());

INSERT INTO tasks (title, description, completed, created_at, updated_at)
VALUES ('Buy groceries', 'Milk, eggs, bread, and fruits', false, NOW(), NOW());

INSERT INTO tasks (title, description, due_date, completed, created_at, updated_at)
VALUES ('Prepare presentation', 'Create slides for the meeting', '2023-07-20 10:30:00', true, NOW(), NOW());
