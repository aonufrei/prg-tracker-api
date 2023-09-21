package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	AuthHeader = "Authorization"
)

func Authorize(c *gin.Context) (*string, error) {
	token := c.GetHeader(AuthHeader)
	if token == "" {
		return nil, errors.New("no authorization token provided")
	}
	return DecodeAuthToken(token)
}

func DecodeAuthToken(token string) (*string, error) {
	parts := strings.Split(token, " ")
	if len(parts) != 2 {
		return nil, errors.New("bad authorization token provided")
	}
	return &parts[1], nil
}
