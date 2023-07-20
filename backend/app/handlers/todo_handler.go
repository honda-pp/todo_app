package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/todo_app/backend/app/usecases"
)

type TodoHandler struct {
	todoUsecase *usecases.TodoUsecase
}

func NewTodoHandler(todoUsecase *usecases.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"todos": []string{"todo1", "todo2"}})
}
