package UserCRUD

import (
	"github.com/gin-gonic/gin"

	Database "MainPackage/Database"
)

type User struct {
	ID   int
	Name string
}

func CreateUser(c *gin.Context) {

	db := Database.ConnectDB()

	name := c.PostForm("name")

	u := User{}
	db.Last(&u)

	user := User{
		ID:   u.ID + 1,
		Name: name,
	}

	db.Create(&user)
	c.JSON(200, gin.H{"name": name})
}

func GetUser(c *gin.Context) {

	db := Database.ConnectDB()

	id := c.Param("id")

	u := User{}
	db.Find(&u, id)

	c.JSON(200, gin.H{"name": u.Name})

}

func GetAllUsers(c *gin.Context) {

	db := Database.ConnectDB()

	allUsers := []User{}
	db.Find(&allUsers)

	for i := 0; i < len(allUsers); i++ {
		c.JSON(200, gin.H{"user": allUsers[i]})
	}

}
