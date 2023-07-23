CREATE TABLE todo (
    task_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date DATE,
    completed boolean NOT NULL DEFAULT false,
    created_at DATE NOT NULL DEFAULT NOW(),
    updated_at DATE
);

INSERT INTO todo (title, description, due_date, completed, created_at)
VALUES ('Finish report', 'Write a report for the project', '2023-07-25 15:00:00', false, NOW());

INSERT INTO todo (title, description, completed, created_at)
VALUES ('Buy groceries', 'Milk, eggs, bread, and fruits', false, NOW());

INSERT INTO todo (title, description, due_date, completed, created_at)
VALUES ('Prepare presentation', 'Create slides for the meeting', '2023-07-20 10:30:00', true, NOW());
