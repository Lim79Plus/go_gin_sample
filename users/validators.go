package users

import (
	"fmt"

	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/gin-gonic/gin"
)

// UserModelValidator struct
type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
		Bio      string `form:"bio" json:"bio" binding:"max=1024"`
		Image    string `form:"image" json:"image" binding:"omitempty,url"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

// NewUserModelValidator return validator
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

// Bind UserModelValidator
func (validate *UserModelValidator) Bind(c *gin.Context) error {
	fmt.Println("UserModelValidator Bind")
	err := common.Bind(c, validate)
	if err != nil {
		fmt.Println("UserModelValidator Bind err", err)
		return err
	}
	validate.userModel.Username = validate.User.Username
	validate.userModel.Email = validate.User.Email
	validate.userModel.Bio = validate.User.Bio

	if validate.User.Password != common.GetNB().NBRandomPassword {
		validate.userModel.setPassword(validate.User.Password)
	}
	if validate.User.Image != "" {
		validate.userModel.Image = &validate.User.Image
	}
	return nil
}

// LoginValidator validate users login request
type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

// NewLoginValidator return loginValidator struct
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}

// Bind for LoginValidator
func (validator *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, validator)
	if err != nil {
		return err
	}

	validator.userModel.Email = validator.User.Email
	return nil
}
