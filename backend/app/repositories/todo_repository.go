package repositories

import (
	"database/sql"
	"time"

	"github.com/honda-pp/todo_app/backend/app/models"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) GetTodoList() ([]*models.Todo, error) {

	query := "SELECT task_id, title, description, due_date, completed, created_at, updated_at FROM todo"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todoList := make([]*models.Todo, 0)

	for rows.Next() {
		todo := &models.Todo{}
		err := rows.Scan(
			&todo.TaskID,
			&todo.Title,
			&todo.Description,
			&todo.DueDate,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todoList = append(todoList, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todoList, nil
}

func (r *TodoRepository) GetTodo(taskID int) (*models.Todo, error) {
	query := "SELECT task_id, title, description, due_date, completed, created_at, updated_at FROM todo WHERE task_id = $1"

	row := r.db.QueryRow(query, taskID)
	todo := &models.Todo{}

	err := row.Scan(
		&todo.TaskID,
		&todo.Title,
		&todo.Description,
		&todo.DueDate,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) CreateTodo(newTodo *models.Todo) (int, error) {
	insertQuery := "INSERT INTO todo (title, description, due_date, completed, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING task_id"
	var taskID int
	err := r.db.QueryRow(insertQuery, newTodo.Title, newTodo.Description, newTodo.DueDate, newTodo.Completed, time.Now()).Scan(&taskID)
	if err != nil {
		return -1, err
	}
	return taskID, nil
}

func (r *TodoRepository) UpdateTodo(todo *models.Todo) error {
	updateQuery := "UPDATE todo SET title = $1, description = $2, due_date = $3, completed = $4, updated_at = $5 WHERE task_id = $6"
	_, err := r.db.Exec(updateQuery, todo.Title, todo.Description, todo.DueDate, todo.Completed, time.Now(), todo.TaskID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) DeleteTodo(taskID int) error {
	deleteQuery := "DELETE FROM todo WHERE task_id = $1"
	_, err := r.db.Exec(deleteQuery, taskID)
	if err != nil {
		return err
	}
	return nil
}
