package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//server and DB initialization
func main(){
	initDB()

	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks", postTask)

	router.Run("localhost:8080")
	fmt.Println("server running on: http://localhost:8080")
}