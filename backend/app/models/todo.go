package models

import (
	"database/sql"
)

type Todo struct {
	TaskID      int          `json:"task_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	DueDate     sql.NullTime `json:"due_date"`
	Completed   bool         `json:"completed"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}
