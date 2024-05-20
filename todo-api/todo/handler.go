package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/todos", GetTodosHandler)
	r.GET("/todos/:id", GetTodoByIDHandler)
	r.POST("/todos", CreateTodoHandler)
	r.PUT("/todos/:id", UpdateTodoHandler)
	r.DELETE("/todos/:id", DeleteTodoHandler)
}

func GetTodosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, GetAllTodos())
}

func GetTodoByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	todo := GetSingleTodoByID(id)
	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func CreateTodoHandler(c *gin.Context) {
	var newTodo struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := CreateNewTodo(newTodo.Title)
	c.JSON(http.StatusCreated, todo)
}

func UpdateTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedTodo struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := UpdateExistingTodoByID(id, updatedTodo.Title, updatedTodo.Completed)
	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if DeleteTodoByID(id) {
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	}
}
