package users

import (
	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/Lim79Plus/go_gin_sample/logger"
	"github.com/gin-gonic/gin"
)

// UserSerializer return gin.Context
type UserSerializer struct {
	c *gin.Context
}

// UserResponse user
type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

// Response return user struct
func (serialize *UserSerializer) Response() UserResponse {
	myUserModel := serialize.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Token:    common.GenToken(myUserModel.ID),
	}
	logger.Trace("UserSerializer.Response()", user)
	return user
}
