package usecases

import (
	"github.com/honda-pp/todo_app/backend/app/models"
	"github.com/honda-pp/todo_app/backend/app/repositories"
)

type TodoUsecase struct {
	todoRepo *repositories.TodoRepository
}

func NewTodoUsecase(todoRepo *repositories.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepo: todoRepo,
	}
}

func (uc *TodoUsecase) GetTodoList() ([]*models.Todo, error) {
	return uc.todoRepo.GetTodoList()
}

func (uc *TodoUsecase) GetTodo(taskID int) (*models.Todo, error) {
	return uc.todoRepo.GetTodo(taskID)
}

func (uc *TodoUsecase) CreateTodo(todo *models.Todo) (int, error) {
	return uc.todoRepo.CreateTodo(todo)
}

func (uc *TodoUsecase) UpdateTodo(todo *models.Todo) error {
	return uc.todoRepo.UpdateTodo(todo)
}

func (uc *TodoUsecase) DeleteTodo(taskID int) error {
	return uc.todoRepo.DeleteTodo(taskID)
}
