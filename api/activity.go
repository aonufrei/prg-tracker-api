package api

import (
	"github.com/gin-gonic/gin"
	"prg-tracker/data"
	"prg-tracker/services"
)

type ActivityApi struct {
	Service *services.ActivityService
}

func (ac *ActivityApi) GetActivity(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityId := c.Param("activityId")
	if activityId == "" {
		c.JSON(400, data.BasicResponse{Message: "Activity was not provided"})
		return
	}
	activityById := ac.Service.GetById(activityId)
	if activityById == nil {
		c.JSON(404, data.BasicResponse{Message: "Activity was not found"})
		return
	}
	if activityById.UserId != *userId {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	c.JSON(200, services.ActivityModelToDto(activityById))
}

func (ac *ActivityApi) GetAllActivities(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activities := make([]data.Activity, 0)
	activities = ac.Service.GetAll(*userId)
	c.JSON(200, services.ActivityModelsToDtos(activities))
}

func (ac *ActivityApi) CreateActivity(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	var activityDto data.ActivityInDto
	if err := c.ShouldBindJSON(&activityDto); err != nil {
		c.JSON(400, data.BasicResponse{Message: "Malformed request body " + err.Error()})
		return
	}
	if !ac.Service.IsNameUnique(*userId, activityDto.Name) {
		c.JSON(400, data.BasicResponse{Message: "Activity with provided name already exists"})
		return
	}
	if err := services.ValidateActivityType(activityDto.Type); err != nil {
		c.JSON(400, data.BasicResponse{Message: err.Error()})
		return
	}
	createdActivity := ac.Service.Create(*userId, data.Activity{
		Name: activityDto.Name,
		Type: activityDto.Type,
	})

	c.JSON(200, services.ActivityModelToDto(createdActivity))
}

func (ac *ActivityApi) UpdateActivity(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityId := c.Param("activityId")
	if activityId == "" {
		c.JSON(400, data.BasicResponse{Message: "Activity was not provided"})
		return
	}
	var activityDto data.ActivityInDto
	if err := c.ShouldBindJSON(&activityDto); err != nil {
		c.JSON(400, data.BasicResponse{Message: "Malformed request body. " + err.Error()})
		return
	}
	if !ac.Service.DoesUserOwnActivity(*userId, activityId) {
		c.JSON(401, data.BasicResponse{Message: "User does not owns activity"})
		return
	}
	if !ac.Service.IsNameUniqueForUpdate(activityId, *userId, activityDto.Name) {
		c.JSON(400, data.BasicResponse{Message: "Activity with provided name already exists"})
		return
	}
	if err := services.ValidateActivityType(activityDto.Type); err != nil {
		c.JSON(400, data.BasicResponse{Message: err.Error()})
		return
	}
	update, updateError := ac.Service.Update(activityId, activityDto)
	if updateError != nil {
		c.JSON(400, data.BasicResponse{Message: "Failed to update activity" + updateError.Error()})
		return
	}
	c.JSON(200, services.ActivityModelToDto(update))
}

func (ac *ActivityApi) DeleteActivity(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityId := c.Param("activityId")
	if activityId == "" {
		c.JSON(400, data.BasicResponse{Message: "Activity was not provided"})
		return
	}
	if !ac.Service.DoesUserOwnActivity(*userId, activityId) {
		c.JSON(401, data.BasicResponse{Message: "User does not owns activity"})
		return
	}
	response := ac.Service.Delete(activityId)
	c.JSON(200, response)
}
