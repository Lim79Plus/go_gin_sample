package articles

import (
	"fmt"
	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

// ArticleModelValidator validator
type ArticleModelValidator struct {
	Article struct {
		Title       string `form:"title" json:"title" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Body        string `form:"body" json:"body" binding:"required"`
	} `json:"article"`
	articleModel ArticleModel `json:"-"`
}

// NewArticleModelValidator return validator struct
func NewArticleModelValidator() ArticleModelValidator {
	return ArticleModelValidator{}
}

// Bind bind article
func (s *ArticleModelValidator) Bind(c *gin.Context) error {
	// myUserModel := c.MustGet("my_user_model").(users.UserModel)
	fmt.Println("ArticleModelValidator Bind")
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.articleModel.Slug = slug.Make(s.Article.Title)
	s.articleModel.Title = s.Article.Title
	s.articleModel.Description = s.Article.Description
	s.articleModel.Body = s.Article.Body
	// s.articleModel.Author = GetArticleUserModel(myUserModel)
	// s.articleModel.setTags(s.Article.Tags)
	return nil
}
