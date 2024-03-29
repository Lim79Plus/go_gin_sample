package main

import (
	"github.com/Lim79Plus/go_gin_sample/articles"
	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/Lim79Plus/go_gin_sample/logger"
	"github.com/Lim79Plus/go_gin_sample/users"
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

// Init set initialization
func Init() {
	// logger setting
	logger.LogInit()

	// config setting
	common.InitConf()

}

func main() {

	Init()

	// db setting
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	// testPage := r.Group("/hello")
	// hello.HelloWorld(testPage.Group("/world"))

	v1 := r.Group("/api")
	users.Register(v1.Group("/register"))
	users.Login(v1.Group("/login"))
	v1.Use(users.AuthMiddleware(false))
	articles.AnonymousRegister(v1.Group("/articles"))

	// following route
	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	articles.Register(v1.Group("/articles"))

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
