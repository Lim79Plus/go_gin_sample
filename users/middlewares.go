package users

import (
	"net/http"
	"strings"

	"github.com/Lim79Plus/go_gin_sample/common"
	"github.com/Lim79Plus/go_gin_sample/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// AuthorizationHeaderExtractor Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// MyAuth2Extractor Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

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

// AuthMiddleware You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//  r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		// Extract jwt token from header
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.GetNB().NBSecretPassword))
			return b, nil
		})
		logger.Trace("AuthMiddleware token:", token)
		if err != nil {
			// the page required auth will be return uauth err when jwt token err
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		// token.Valid request.ParseFromRequest return the result.
		// token.Claims.(jwt.MapClaims) => cast accertion
		logger.Trace("AuthMiddleware token:", token.Claims)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			myUserID := uint(claims["id"].(float64))
			UpdateContextUserModel(c, myUserID)
		}
	}
}
