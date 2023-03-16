package main

import (
	"log"

	"github.com/gin-gonic/gin"

	UserCRUD "MainPackage/Api"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	api := gin.Default()

	api.POST("/post", UserCRUD.CreatePost)
	api.GET("/userinfo/:id", UserCRUD.GetPost)
	api.GET("/userlist", UserCRUD.GetAllPosts)

	log.Printf("Server started successfully")

	api.Run()

}
