package users

import "github.com/gin-gonic/gin"

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
		// Token:    common.GenToken(myUserModel.ID),
	}
	return user
}
