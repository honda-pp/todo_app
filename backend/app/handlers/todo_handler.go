package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/todo_app/backend/app/models"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todo list"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"todoList": todoList})
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	taskID, err := h.todoUsecase.CreateTodo(&newTodo)
	if err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully", "taskId": taskID})
}
