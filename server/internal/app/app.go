package app

import (
	"log"
	"todo-list/internal/handler"
	"todo-list/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func App() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	database.Connect()
	router := gin.Default()

	router.GET("/", handler.GetAll)
	router.GET("/:id", handler.FindById)

	router.Run(":8787")
}
