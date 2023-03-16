package Database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	Models "MainPackage/Models"
)

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

	db.AutoMigrate(&Models.UserAuth{})
	log.Println("Database Migration Completed!")

	return db
}
