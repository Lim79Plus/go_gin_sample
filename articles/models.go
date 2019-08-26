package articles

import "github.com/jinzhu/gorm"
import "github.com/Lim79Plus/go_gin_sample/common"

// ArticleModel struct
type ArticleModel struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	// Author      ArticleUserModel
	// AuthorID    uint
	// Tags        []TagModel     `gorm:"many2many:article_tags;"`
	// Comments    []CommentModel `gorm:"ForeignKey:ArticleID"`
}

// FindArticleList return list of article
func FindArticleList() ([]ArticleModel, int, error) {
	db := common.GetDB()
	var models []ArticleModel
	var count int

	offsetInt := 0
	limitInt := 20

	tx := db.Begin()
	db.Model(&models).Count(&count)
	db.Offset(offsetInt).Limit(limitInt).Find(&models)

	err := tx.Commit().Error
	return models, count, err
}

// FindOneArticle return one article
func FindOneArticle(condition interface{}) (ArticleModel, error) {
	db := common.GetDB()
	var model ArticleModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	// tx.Model(&model).Related(&model.Author, "Author")
	// tx.Model(&model.Author).Related(&model.Author.UserModel)
	// tx.Model(&model).Related(&model.Tags, "Tags")
	err := tx.Commit().Error
	return model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}