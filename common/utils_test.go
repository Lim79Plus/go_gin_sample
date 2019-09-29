package common

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lim79Plus/go_gin_sample/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func initConf() {
	ConfName = "../config.ini"
	InitConf()
	logger.LogInit()
}
func TestGenToken(t *testing.T) {
	initConf()
	id := uint(1234)
	token := GenToken(id)
	t.Log("TestGenToken result:", token)
	if token == "" {
		t.Fatal("failed test", "result", token)
	}
	if len(token) != 119 {
		t.Fatal("failed to not validate length", "result", len(token))
	}
}

func TestNewValidatorError(t *testing.T) {
	asserts := assert.New(t)
	initConf()

	type Login struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	}

	var requestTest = struct {
		bodyData       string
		expectedCode   int
		responseRegexg string
		msg            string
	}{
		`{"username": "wangzitian0","password": "0122"}`,
		http.StatusUnprocessableEntity,
		`{"errors":{"Password":"{min: 8}"}}`,
		"invalid password of too short and should return StatusUnprocessableEntity",
	}

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var json Login
		if err := Bind(c, &json); err != nil {
			c.JSON(http.StatusUnprocessableEntity, NewValidatorError(err))
		}
	})

	testData := requestTest
	bodyData := testData.bodyData
	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(bodyData))
	req.Header.Set("Content-Type", "application/json")
	asserts.NoError(err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
	asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
}

func TestNewError(t *testing.T) {
	assert := assert.New(t)
	key := "validate"
	errMsg := "err message"
	res := NewError(key, errors.New(errMsg))
	assert.IsType(res, res, "commonerr dose not return right type")
	assert.Equal(map[string]interface{}(map[string]interface{}{key: errMsg}),
		res.Errors, "commenError should have right error info")
}
