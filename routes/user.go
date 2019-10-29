package routes

import (
	"github.com/Mbenx/Rest-Api-Basic/config"
	"github.com/Mbenx/Rest-Api-Basic/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	items := []models.Karyawan{}
	// db.First(&article, 1)
	config.DB.Find(&items)

	c.JSON(200, gin.H{
		"message": "Data Karyawan",
		"data":    items,
	})
}

func PostUsers(c *gin.Context) {

}
