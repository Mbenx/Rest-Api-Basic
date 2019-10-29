package config

import (
	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // for mysql database
)

// DB ...
var DB *gorm.DB

// InitDB ...
func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:h3ru@mysql@/php_basic")
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(models.Article{})
}
