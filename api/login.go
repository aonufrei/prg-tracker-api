package api

import (
	"github.com/gin-gonic/gin"
	"prg-tracker/data"
	"prg-tracker/services"
)

type LoginApi struct {
	UserService *services.UserService
}

func (la *LoginApi) Login(c *gin.Context) {
	loginRequest := data.LoginInDto{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, data.BasicResponse{Message: "Malformed request"})
		return
	}
	if loginRequest.Username == nil && loginRequest.Password == nil {
		c.JSON(400, data.BasicResponse{Message: "Both username and password is required"})
		return
	}
	request, err := services.ProcessLoginRequest(loginRequest, la.UserService)
	if err != nil {
		c.JSON(401, data.BasicResponse{Message: err.Error()})
		return
	}
	c.JSON(200, data.LoginOutDto{Token: request})
}
