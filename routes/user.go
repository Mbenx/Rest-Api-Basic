package routes

import (
	"../config"
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	items := models.[]Karyawan{}
	// db.First(&article, 1)
	config.DB.Find(&items)

	c.JSON(200, gin.H{
		"message": "Data Karyawan",
		"data":    items,
	})
}

func postUsers(c *gin.Context) {

}
