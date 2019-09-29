package articles

import (
	"net/http"

	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/Lim79Plus/go_gin_sample/logger"
	"github.com/gin-gonic/gin"
)

// Register for login user
func Register(router *gin.RouterGroup) {
	router.POST("/", ArticleCreate)
}

// AnonymousRegister register public route for article
func AnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", ArticleList)
	router.GET("/:slug", ArticleRetrieve)
}

// ArticleCreate register
func ArticleCreate(c *gin.Context) {
	logger.Trace("ArticleCreate start", c.Request.Method, c.ContentType())

	articleModelValidator := NewArticleModelValidator()
	if err := articleModelValidator.Bind(c); err != nil {
		logger.Trace("ArticleCreate err happend", err)
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	logger.Trace("success", articleModelValidator)
	//logger.Trace(articleModelValidator.articleModel.Author.UserModel)

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
	logger.Trace("ArticleList serializer", serializer)
	c.JSON(http.StatusOK, gin.H{"articles": serializer.Response(), "articlesCount": modelCount})
}

// ArticleRetrieve retreve one article
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
