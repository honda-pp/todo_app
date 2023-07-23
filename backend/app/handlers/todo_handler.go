package handlers

import (
	"net/http"
	"strconv"

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

func (h *TodoHandler) GetTodo(c *gin.Context) {

	taskIDStr := c.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}

	todo, err := h.todoUsecase.GetTodo(taskID)
	if err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todo"})
		return
	}

	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	taskID, err := h.todoUsecase.CreateTodo(&newTodo)
	if err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task_id": taskID})
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var updatedTodo models.Todo
	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	taskIDStr := c.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}

	if taskID != updatedTodo.TaskID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_id in URL does not match task_id in request body"})
		return
	}

	err = h.todoUsecase.UpdateTodo(&updatedTodo)
	if err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	taskIDStr := c.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}

	err = h.todoUsecase.DeleteTodo(taskID)
	if err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
