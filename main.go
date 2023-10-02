package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string
	Item      string
	Completed bool
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(contex *gin.Context) {
	contex.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusOK, newTodo)
}

func main() {
	router := gin.Default() // The server
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
