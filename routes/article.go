package routes

import (
	"github.com/Mbenx/Rest-Api-Basic/config"
	"github.com/Mbenx/Rest-Api-Basic/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

// GetArticle ...
func GetArticle(c *gin.Context) {
	slug := c.Param("slug")
	// item := []models.Article{}
	var item models.Article
	if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	// log.Printf("Hasil Query", item)

	c.JSON(200, gin.H{
		"message": "Welcome to Golang",
		"slug":    slug,
		"data":    item,
	})
}

// PostArticle ...
func PostArticle(c *gin.Context) {
	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}
	// Create
	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"message": "Post Berhasil",
		"data":    item,
	})
}
