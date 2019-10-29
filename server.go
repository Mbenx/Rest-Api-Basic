package main

import (
	"github.com/Mbenx/Rest-Api-Basic/config"
	"github.com/Mbenx/Rest-Api-Basic/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", getHome)
		v1.GET("/article/:slug", routes.GetArticle)
		v1.POST("/article", routes.PostArticle)

		blogs := v1.Group("/blogs")
		{
			blogs.GET("/", routes.GetBlogs)
			blogs.POST("/", routes.PostBlogs)
		}

		users := v1.Group("/users")
		{
			users.GET("/", routes.GetUsers)
			users.POST("/", routes.PostUsers)
		}
	}

	router.Run()
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Golang",
	})
}
