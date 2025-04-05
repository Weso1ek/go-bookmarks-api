package main

import (
	"bookmark-api/context"
	"bookmark-api/httpcontroller"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	config := context.GetConfig()

	fmt.Println("===")
	fmt.Println(config.DBUser)
	fmt.Println(config.DBPassword)
	fmt.Println(config.DBHost)
	fmt.Println("===")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}

	log.Println("Server start")

	categoriesController := httpcontroller.NewCategoriesController()

	router := gin.Default()
	router.GET("/categories", categoriesController.GetCategories)

	router.Run("localhost:8080")
}
