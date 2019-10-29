package routes

import (
	"../config"
	"../models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetArticle(c *gin.Context) {
	slug := c.Param("slug")
	// items := []Article{}
	var items models.Article
	if config.DB.First(&items, "slug = ?", slug).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Welcome to Golang",
		"slug":    slug,
		"data":    items,
	})
}

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
