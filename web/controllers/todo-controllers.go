package controllers

import (
	"awesomeProject/config"
	"awesomeProject/lib"
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Compeleted  bool   `json:"compeleted"`
}

type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

func CreateTodo(context *gin.Context) {
	var data todoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func Test(context *gin.Context) {

	bearerToken := context.Request.Header.Get("Authorization")
	if bearerToken == "" || bearerToken == "null" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Bearer token is required",
		})
	}

	var authorizationToken = strings.Split(bearerToken, " ")[1]

	email, _ := lib.VerifyToken(authorizationToken)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    email,
	})

}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// var response []todoResponse
	// for _, todo := range todos {
	// 	response = append(response, todoResponse{
	// 		ID:          todo.ID,
	// 		name:        todo.Name,
	// 		description: todo.Description,
	// 	})
	// }

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})

}

func GetTodo(context *gin.Context) {
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	var todo models.Todo

	err := db.Where("id = ?", idTodo).First(&todo)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todo,
	})

}

func UpdateTodo(context *gin.Context) {
	var data todoRequest
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}

	todoById := db.Where("id = ?", idTodo).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func CompleteTodo(context *gin.Context) {
	//var data todoRequest
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	//if err := context.BindJSON(&data); err != nil {
	//	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	todo := models.Todo{}

	todoById := db.Where("id = ?", idTodo).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	todo.Completed = true

	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description
	response.Compeleted = todo.Completed

	context.JSON(http.StatusCreated, response)
}

func DeleteTodo(context *gin.Context) {
	todo := models.Todo{}
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", idTodo).Unscoped().Delete(&todo)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idTodo,
	})

}
