package usecases

import "github.com/honda-pp/todo_app/backend/app/repositories"

type TodoUsecase struct {
	todoRepo *repositories.TodoRepository
}

func NewTodoUsecase(todoRepo *repositories.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepo: todoRepo,
	}
}
