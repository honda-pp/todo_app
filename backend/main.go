package main

import (
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:88"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	todoHandler := handlers.NewTodoHandler(todoUsecase)

	api := router.Group("/api")
	{
		api.GET("/todoList", todoHandler.GetTodoList)
		api.POST("/todo", todoHandler.CreateTodo)
		api.PUT("/todoList/:task_id", todoHandler.UpdateTodo)
		api.DELETE("/todoList/:task_id", todoHandler.DeleteTodo)
	}

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
