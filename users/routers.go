package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/gin-gonic/gin"
)

// Register regist new user
func Register(router *gin.RouterGroup) {
	router.POST("/", Registration)
}

// Login to the site
func Login(router *gin.RouterGroup) {
	router.POST("/", login)
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

func login(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := findOneUser(&UserModel{Email: loginValidator.userModel.Email})
	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	UpdateContextUserModel(c, userModel.ID)
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
