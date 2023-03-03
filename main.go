package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	ConnectDB()

	gin.SetMode(gin.ReleaseMode)
	api := gin.Default()

	api.POST("/post", func(c *gin.Context) {
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.JSON(200, gin.H{"name": name, "message": message})
	})
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Printf("Server started successfully")
	api.Run(":3000")
}

func ConnectDB() *gorm.DB {
	err := godotenv.Load("Database.env")
	if err != nil {
		log.Fatal("Database.env load failed")
	} else {
		log.Printf("Database.env loaded successfully")
	}

	db_host := os.Getenv("DB_Host")
	db_port := os.Getenv("DB_Port")
	db_database := os.Getenv("DB_Database")
	db_username := os.Getenv("DB_Username")
	db_password := os.Getenv("DB_Password")

	dsn := "host= " + db_host + " user=" + db_username + " password=" + db_password + " dbname=" + db_database + " port=" + db_port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
	} else {
		log.Printf("Database connection succeeded")
	}

	return db
}
