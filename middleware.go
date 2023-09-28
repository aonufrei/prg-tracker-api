package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"prg-tracker/data"
	"prg-tracker/services"
)

func createCorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	})
}

func getUserFromRequest(c *gin.Context, userService services.UserService) (*data.User, error) {
	token := c.GetHeader(services.AuthHeader)
	userId, authError := services.DecodeAuthToken(token)
	if authError != nil {
		return nil, authError
	}
	user := userService.GetById(*userId)
	return user, nil
}

func requireSpecificRole(requiredRole string, userService services.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		user, authError := getUserFromRequest(c, userService)
		if authError != nil || user == nil || user.Role != requiredRole {
			//c.JSON()
			c.AbortWithStatusJSON(401, data.BasicResponse{Message: "Unauthorized"})
			return
		}
		c.Next()
	}
}
