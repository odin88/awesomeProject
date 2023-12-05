package routes

import (
	controllers2 "awesomeProject/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers2.CreateTodo)
	route.GET("/todo", controllers2.GetAllTodos)
	route.GET("/test", controllers2.Test)
	route.GET("/todo/:idTodo", controllers2.GetTodo)
	route.PUT("/todo/:idTodo", controllers2.UpdateTodo)
	route.PUT("/todo/complete/:idTodo", controllers2.CompleteTodo)
	route.DELETE("/todo/:idTodo", controllers2.DeleteTodo)

	route.GET("/facebook/login", controllers2.FacebookLogin)
	route.GET("/facebook/callback", controllers2.HandleFacebookLogin)

	err := route.Run()

	if err != nil {
		log.Fatal("Error occcured")
	}
}
