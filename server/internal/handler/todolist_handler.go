package handler

import (
	"net/http"
	"todo-list/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	todos, err := repository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to retrieve todo",
			"details": err.Error(),
		})
		return
	}

	if todos.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "No todo found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    todos,
	})
}

func FindById(c *gin.Context) {
	todoId := c.Params.ByName("id")
	todos, err := repository.FindById(todoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to retrieve todo",
			"details": err.Error(),
		})
		return
	}

	if todos.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "No todo found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    todos,
	})
}
