package users

import (
	"fmt"
	"net/http"

	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/gin-gonic/gin"
)

// Register regist new user
func Register(router *gin.RouterGroup) {
	router.POST("/", Registration)
}

// Registration create new user account
func Registration(c *gin.Context) {
	fmt.Println("user Registration start", c.Request.Method, c.ContentType())
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		fmt.Println("user Registration err happend", err)
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
