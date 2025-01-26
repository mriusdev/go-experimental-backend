package main

import (
	"api-backend/sample/app/models"
	"api-backend/sample/handler"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlServerIp := os.Getenv("MYSQL_SERVER_IP")
	mysqlPort := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		mysqlUser, mysqlUserPassword, mysqlServerIp, mysqlPort, mysqlDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed connecting to database")
		fmt.Println(err.Error())
		return
	}
	db.AutoMigrate(&models.User{}, &models.Permission{}, &models.Group{}, &models.Post{})

	handler.HandleRequest()
}
