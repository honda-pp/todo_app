package main

import (
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/todo_app/backend/app/handlers"
	"github.com/honda-pp/todo_app/backend/app/repositories"
	"github.com/honda-pp/todo_app/backend/app/usecases"
	"github.com/honda-pp/todo_app/backend/infrastructure/database"
	"github.com/honda-pp/todo_app/backend/infrastructure/logger"
)

func main() {
	log := logger.InitLogger()
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	todoRepo := repositories.NewTodoRepository(db)

	todoUsecase := usecases.NewTodoUsecase(todoRepo)

	router := gin.Default()

	todoHandler := handlers.NewTodoHandler(todoUsecase)

	api := router.Group("/api")
	{
		api.GET("/todos", todoHandler.GetTodos)
		/*
			api.POST("/todos", todoHandler.CreateTodo)
			api.PUT("/todos/:id", todoHandler.UpdateTodo)
			api.DELETE("/todos/:id", todoHandler.DeleteTodo)
		*/
	}

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
