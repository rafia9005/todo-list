package app

import "github.com/gin-gonic/gin"

func App() {
	router := gin.Default()
	router.Run(":8787")
}
