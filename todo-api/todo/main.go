package main

import (
	"github.com/Ashikurrahaman287/todo-api/todo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	todo.RegisterRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
