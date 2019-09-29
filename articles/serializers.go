package articles

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type ArticleSerializer struct {
	C *gin.Context
	ArticleModel
}

type ArticleListResponse struct {
	ID          uint   `json:"-"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Body        string `json:"body"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	// Author         users.ProfileResponse `json:"author"`
	// Tags           []string              `json:"tagList"`
	// Favorite       bool                  `json:"favorited"`
	// FavoritesCount uint                  `json:"favoritesCount"`
}

type ArticleListSerializer struct {
	C        *gin.Context
	Articles []ArticleModel
}

func (s *ArticleSerializer) Response() ArticleListResponse {
	// myUserModel := s.C.MustGet("my_user_model").(users.UserModel)
	// authorSerializer := ArticleUserSerializer{s.C, s.Author}
	response := ArticleListResponse{
		ID:          s.ID,
		Slug:        slug.Make(s.Title),
		Title:       s.Title,
		Description: s.Description,
		Body:        s.Body,
		CreatedAt:   s.CreatedAt.UTC().Format(time.RFC3339),
		// CreatedAt: s.CreatedAt.String(),
		UpdatedAt: s.UpdatedAt.UTC().Format(time.RFC3339),
		// UpdatedAt: s.UpdatedAt.String(),
		// Author:         authorSerializer.Response(),
		// Favorite:       s.isFavoriteBy(GetArticleUserModel(myUserModel)),
		// FavoritesCount: s.favoritesCount(),
	}
	// response.Tags = make([]string, 0)
	// for _, tag := range s.Tags {
	// 	serializer := TagSerializer{s.C, tag}
	// 	response.Tags = append(response.Tags, serializer.Response())
	// }
	return response
}

// Response for Articles
func (s *ArticleListSerializer) Response() []ArticleListResponse {
	response := []ArticleListResponse{}
	for _, article := range s.Articles {
		serializer := ArticleSerializer{s.C, article}
		response = append(response, serializer.Response())
	}
	return response
}
