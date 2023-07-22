package repositories

import (
	"database/sql"

	"github.com/honda-pp/todo_app/backend/app/models"
	"github.com/honda-pp/todo_app/backend/infrastructure/logger"
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
		logger.LogInfo("ddd")
		todoList = append(todoList, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todoList, nil
}
