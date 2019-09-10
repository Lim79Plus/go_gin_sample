package users

import (
	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/gin-gonic/gin"
)

// UpdateContextUserModel is a helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, userID uint) {
	var myUserModel UserModel
	if userID != 0 {
		db := common.GetDB()
		db.First(&myUserModel, userID)
	}
	// Set is used to store a new key/value pair exclusively for this context.
	c.Set("my_user_id", userID)
	c.Set("my_user_model", myUserModel)
}
