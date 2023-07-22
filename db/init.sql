CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMP,
    status VARCHAR(20) NOT NULL DEFAULT '未完了',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

INSERT INTO tasks (title, description, due_date, status, created_at, updated_at)
VALUES ('Finish report', 'Write a report for the project', '2023-07-25 15:00:00', 'In progress', NOW(), NOW());

INSERT INTO tasks (title, description, status, created_at, updated_at)
VALUES ('Buy groceries', 'Milk, eggs, bread, and fruits', 'Not started', NOW(), NOW());

INSERT INTO tasks (title, description, due_date, status, created_at, updated_at)
VALUES ('Prepare presentation', 'Create slides for the meeting', '2023-07-20 10:30:00', 'Completed', NOW(), NOW());
