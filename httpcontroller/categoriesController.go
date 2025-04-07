package httpcontroller

import (
	"bookmark-api/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type CategoriesController struct {
	db *gorm.DB
}

func NewCategoriesController(db *gorm.DB) *CategoriesController {
	return &CategoriesController{
		db: db,
	}
}

func (cc CategoriesController) GetCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func (cc CategoriesController) PostCategories(c *gin.Context) {
	var newCategory dto.Category

	if err := c.ShouldBindJSON(&newCategory); err != nil {

		fmt.Println("====")
		fmt.Println(err)
		fmt.Println("====")

		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	errSave := cc.db.Save(&newCategory).Error

	if errSave != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errSave.Error()})
	}

	fmt.Println(newCategory)

}
