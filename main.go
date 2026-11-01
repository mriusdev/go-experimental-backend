package main

import (
	"api-backend/sample/app/application"
	"api-backend/sample/app/application/env"
	"api-backend/sample/app/application/paths"
	"api-backend/sample/app/models"
	"api-backend/sample/handler"
	"fmt"
	"os"

	"log/slog"

	"github.com/go-chi/chi/v5"
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

	logFilePath := fmt.Sprintf("%s%s", paths.Root, "/var/logs.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	
	logger := slog.New(
		slog.NewJSONHandler(logFile, nil),
	)
	logger.Info("SOmething something", "hello", 13)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed connecting to database")
		fmt.Println(err.Error())
		return
	}
	db.AutoMigrate(&models.User{}, &models.Permission{}, &models.Group{}, &models.Post{})

	api := application.NewApiHandler(chi.NewRouter(), logger, db)
	handler.HandleRequest(api)
}
