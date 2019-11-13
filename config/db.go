package config

import (
	"github.com/Mbenx/Rest-Api-Basic/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // for mysql database
)

// DB ...
var DB *gorm.DB

// InitDB ...
func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:H3ru@mysql@/go_basic")
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Article{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	DB.Model(&models.User{}).Related(&models.Article{})
}
