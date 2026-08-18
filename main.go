package main

import (
	"api-backend/sample/app/application"
	"api-backend/sample/app/application/env"
	"api-backend/sample/app/models"
	"api-backend/sample/handler"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		env.MysqlUser.GetValue(),
		env.MysqlPassword.GetValue(),
		env.MysqlServerIp.GetValue(),
		env.MysqlPort.GetValue(),
		env.MysqlDatabase.GetValue(),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed connecting to database")
		fmt.Println(err.Error())
		return
	}
	application.DB = db
	application.DB.AutoMigrate(&models.User{}, &models.Permission{}, &models.Group{}, &models.Post{})

	handler.HandleRequest()
}
