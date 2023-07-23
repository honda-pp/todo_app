package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/todo_app/backend/app/usecases"
	"github.com/honda-pp/todo_app/backend/infrastructure/logger"
)

type TodoHandler struct {
	todoUsecase *usecases.TodoUsecase
}

func NewTodoHandler(todoUsecase *usecases.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (h *TodoHandler) GetTodoList(c *gin.Context) {

	todoList, err := h.todoUsecase.GetTodoList()
	if err != nil {
		logger.LogError(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"todoList": todoList})
}
