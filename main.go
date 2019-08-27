package main

import (
	"github.com/Lim79Plus/go_gin_sample/articles"
	"github.com/Lim79Plus/go_gin_sample/users"
	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/Lim79Plus/go_gin_sample/hello"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Migrate generate table
func Migrate(db *gorm.DB) {
	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON'T change existing column's type or delete unused columns to protect your data.
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&users.UserModel{})
}

func main() {
	common.InitConf()
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	testPage := r.Group("/hello")
	hello.HelloWorld(testPage.Group("/world"))

	v1 := r.Group("/api")
	articles.AnonymousRegister(v1.Group("/articles"))
	articles.ArticlesRegister(v1.Group("/articles"))

	// start server
	r.Run(common.GetWebPort())
}

// func helloWorld(router *gin.RouterGroup) {
// 	router.GET("/", func(c *gin.Context) {
// 		c.String(200, "Hello World! by group")
// 	})
// }

// func helloWorld(c *gin.Context) {
// 	c.String(200, "Hello World!")
// }
