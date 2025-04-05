package main

import (
	"bookmark-api/httpcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	categoriesController := httpcontroller.NewCategoriesController()

	router := gin.Default()
	router.GET("/albums", categoriesController.GetCategories)

	router.Run("localhost:8080")
}
