package articles

import (
	"fmt"
	"net/http"

	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/gin-gonic/gin"
)

func ArticlesRegister(router *gin.RouterGroup) {
	router.POST("/", ArticleCreate)
}

// AnonymousRegister register public route for article
func AnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", ArticleList)
	router.GET("/:slug", ArticleRetrieve)
}

func ArticleCreate(c *gin.Context) {
	fmt.Println("ArticleCreate start", c.Request.Method)
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	b := string(buf[0:n])
	fmt.Println(b)

	articleModelValidator := NewArticleModelValidator()
	if err := articleModelValidator.Bind(c); err != nil {
		fmt.Println("ArticleCreate err happend", err)
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	//fmt.Println(articleModelValidator.articleModel.Author.UserModel)

	if err := SaveOne(&articleModelValidator.articleModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ArticleSerializer{c, articleModelValidator.articleModel}
	c.JSON(http.StatusCreated, gin.H{"article": serializer.Response()})
}

// ArticleList return json format article list.
func ArticleList(c *gin.Context) {
	//condition := ArticleModel{}
	// tag := c.Query("tag")
	// author := c.Query("author")
	// favorited := c.Query("favorited")
	// limit := c.Query("limit")
	// offset := c.Query("offset")
	// articleModels, modelCount, _ := FindManyArticle(tag, author, limit, offset, favorited)
	articleModels, modelCount, _ := FindArticleList()
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, common.NewError("articles", errors.New("Invalid param")))
	// 	return
	// }
	serializer := ArticleListSerializer{c, articleModels}
	c.JSON(http.StatusOK, gin.H{"articles": serializer.Response(), "articlesCount": modelCount})
}

func ArticleRetrieve(c *gin.Context) {
	slug := c.Param("slug")
	// if slug == "feed" {
	// 	ArticleFeed(c)
	// 	return
	// }
	articleModel, _ := FindOneArticle(&ArticleModel{Slug: slug})
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, common.NewError("articles", errors.New("Invalid slug")))
	// 	return
	// }
	serializer := ArticleSerializer{c, articleModel}
	c.JSON(http.StatusOK, gin.H{"article": serializer.Response()})
}
