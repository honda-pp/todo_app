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

func (r *TodoRepository) CreateTodo(newTodo *models.Todo) (int, error) {
	insertQuery := "INSERT INTO todo (title, description, due_date, completed, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING task_id"
	var taskID int
	err := r.db.QueryRow(insertQuery, newTodo.Title, newTodo.Description, newTodo.DueDate, newTodo.Completed, time.Now()).Scan(&taskID)
	if err != nil {
		return -1, err
	}
	return taskID, nil
}

func (r *TodoRepository) DeleteTodo(taskID int) error {
	deleteQuery := "DELETE FROM todo WHERE task_id = $1"
	_, err := r.db.Exec(deleteQuery, taskID)
	if err != nil {
		return err
	}
	return nil
}
