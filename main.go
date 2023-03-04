package main

import (
	"log"

	"github.com/gin-gonic/gin"

	UserCRUD "MainPackage/Api"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	api := gin.Default()

	api.POST("/post", UserCRUD.CreateUser)
	api.GET("/userinfo/:id", UserCRUD.GetUser)
	api.GET("/userlist", UserCRUD.GetAllUsers)

	log.Printf("Server started successfully")
	api.Run()
}
