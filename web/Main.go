package main

import (
	"awesomeProject/config"
	"awesomeProject/routes"
	"fmt"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	/*r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080*/

	fmt.Println("http://localhost:3333/facebook/login")
	defer config.DisconnectDB(db)
	//run all routes
	routes.Routes()
}
