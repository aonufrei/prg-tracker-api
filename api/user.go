package api

import (
	"github.com/gin-gonic/gin"
	"prg-tracker/data"
	"prg-tracker/services"
)

type UserApi struct {
	Service *services.UserService
}

func (uc *UserApi) GetUser(c *gin.Context) {
	userId := c.Param("id")
	userById := uc.Service.GetById(userId)
	if userById == nil {
		c.JSON(404, data.BasicResponse{Message: "User was not found"})
		return
	}
	c.JSON(200, services.UserModelToDto(userById))
}

func (uc *UserApi) GetAllUsers(c *gin.Context) {
	users := make([]data.User, 0)
	users = uc.Service.GetAll()
	c.JSON(200, services.UserModelsToDtos(users))
}

func (uc *UserApi) CreateUser(c *gin.Context) {
	var userDto data.UserInDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, data.BasicResponse{Message: "Malformed request body " + err.Error()})
		return
	}
	if err := services.ValidateUserInDto(userDto); err != nil {
		c.JSON(400, data.BasicResponse{Message: err.Error()})
		return
	}
	if err := services.ValidateUsername(userDto.Username, uc.Service); err != nil {
		c.JSON(400, data.BasicResponse{Message: err.Error()})
		return
	}
	createdUser := uc.Service.Create(data.User{
		Name:     userDto.Name,
		Username: userDto.Username,
		Password: userDto.Password,
		Role:     userDto.Role,
	})
	c.JSON(200, services.UserModelToDto(createdUser))
}

func (uc *UserApi) UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	var userDto data.UserInDto
	if err := services.ValidateUserIdForUpdate(userId, uc.Service); err != nil {
		c.JSON(400, data.BasicResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, data.BasicResponse{Message: "Malformed request body. " + err.Error()})
		return
	}
	if err := services.ValidateUserInDto(userDto); err != nil {
		c.JSON(400, data.BasicResponse{Message: err.Error()})
		return
	}
	if err := services.ValidateUsernameForUser(userId, userDto.Username, uc.Service); err != nil {
		c.JSON(400, data.BasicResponse{Message: err.Error()})
		return
	}
	update, updateError := uc.Service.Update(userId, userDto)
	if updateError != nil {
		c.JSON(400, data.BasicResponse{Message: "Failed to update user. " + updateError.Error()})
		return
	}
	c.JSON(200, services.UserModelToDto(update))
}

func (uc *UserApi) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	response := uc.Service.Delete(userId)
	c.JSON(200, response)
}
