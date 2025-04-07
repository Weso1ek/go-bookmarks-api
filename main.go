package main

import (
	"bookmark-api/context"
	"bookmark-api/httpcontroller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	config := context.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}

	fmt.Println("Connected to db")
	log.Println("Server start")

	categoriesController := httpcontroller.NewCategoriesController(db)

	router := gin.Default()
	router.GET("/categories", categoriesController.GetCategories)
	router.POST("/categories", categoriesController.PostCategories)

	router.Run(":8080")
}
