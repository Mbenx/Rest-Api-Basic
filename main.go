package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Karyawan struct {
	gorm.Model
	Nama, Alamat, Jabatan string
}

func (Karyawan) TableName() string {
	return "karyawan"
}

type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text"`
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", getHome)
		v1.GET("/article/:slug", getArticle)
		v1.POST("/article", postArticle)

		blogs := v1.Group("/blogs")
		{
			blogs.GET("/", getBlogs)
			blogs.POST("/", postBlogs)
		}

		users := v1.Group("/users")
		{
			users.GET("/", getUsers)
			users.POST("/", postUsers)
		}
	}

	var err error
	DB, err = gorm.Open("mysql", "root:h3ru@mysql@/php_basic")
	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&Article{})

	// Read
	// var article Article
	// db.First(&article, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	// db.Delete(&product)

	// text := slug.Make()

	router.Run() // listen and serve on 0.0.0.0:8080
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Golang",
	})
}

func getArticle(c *gin.Context) {
	slug := c.Param("slug")
	// items := []Article{}
	var items Article
	if DB.First(&items, "slug = ?", slug).RecordNotFound() {
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

func postArticle(c *gin.Context) {
	item := Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}
	// Create
	DB.Create(&item)

	c.JSON(200, gin.H{
		"message": "Post Berhasil",
		"data":    item,
	})
}

func getUsers(c *gin.Context) {
	items := []Karyawan{}
	// db.First(&article, 1)
	DB.Find(&items)

	c.JSON(200, gin.H{
		"message": "Data Karyawan",
		"data":    items,
	})
}

func postUsers(c *gin.Context) {

}

func getBlogs(c *gin.Context) {

}

func postBlogs(c *gin.Context) {

}
