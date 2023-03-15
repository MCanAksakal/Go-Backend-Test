package UserCRUD

import (
	"time"

	"github.com/gin-gonic/gin"

	Database "MainPackage/Database"

	"strconv"
)

type Post struct {
	ID    int
	Label string
	Title string
	Body  string
	Image string
}

func CreatePost(c *gin.Context) {

	db := Database.ConnectDB()

	label := c.PostForm("label")
	title := c.PostForm("title")
	body := c.PostForm("body")
	image := c.PostForm("image")

	u := Post{}
	db.Last(&u)

	post := Post{
		ID:    u.ID,
		Label: label,
		Title: title,
		Body:  body,
		Image: image,
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(200, gin.H{
			"message": "failed",
			"log":     err.Error(),
			"date":    time.Now().Format("2006-01-02 15:04:05"),
			"ip":      c.ClientIP(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
			"log":     "Post created successfully. Post id: " + strconv.Itoa(u.ID+1),
			"date":    time.Now().Format("2006-01-02 15:04:05"),
			"ip":      c.ClientIP(),
		})
	}

}

func GetPost(c *gin.Context) {

	db := Database.ConnectDB()

	id := c.Param("id")

	u := Post{}

	if err := db.Find(&u, id).Error; err != nil {
		c.JSON(200, gin.H{
			"message": "failed",
			"log":     err.Error(),
			"date":    time.Now().Format("2006-01-02 15:04:05"),
			"ip":      c.ClientIP(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
			"log":     "Post readed successfully. Post id: " + strconv.Itoa(u.ID),
			"date":    time.Now().Format("2006-01-02 15:04:05"),
			"ip":      c.ClientIP(),
			"post":    u,
		})
	}

}

func GetAllPosts(c *gin.Context) {

	db := Database.ConnectDB()

	allPosts := []Post{}

	if err := db.Find(&allPosts).Error; err != nil {
		c.JSON(200, gin.H{
			"message": "failed",
			"log":     err.Error(),
			"date":    time.Now().Format("2006-01-02 15:04:05"),
			"ip":      c.ClientIP(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
			"log":     "All posts readed successfully.",
			"date":    time.Now().Format("2006-01-02 15:04:05"),
			"ip":      c.ClientIP(),
			"post":    allPosts,
		})
	}

}
