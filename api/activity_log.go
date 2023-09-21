package api

import (
	"github.com/gin-gonic/gin"
	"prg-tracker/data"
	"prg-tracker/services"
	"time"
)

type ActivityLogApi struct {
	Service         *services.ActivityLogService
	ActivityService *services.ActivityService
}

func (ac *ActivityLogApi) GetActivityLog(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityId, logId := c.Param("activityId"), c.Param("logId")
	if activityId == "" || logId == "" {
		c.JSON(400, data.BasicResponse{Message: "Malformed request"})
		return
	}
	if !ac.ActivityService.DoesUserOwnActivity(*userId, activityId) {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	if !ac.ActivityService.DoesActivityOwnsLog(activityId, logId, ac.Service) {
		c.JSON(404, data.BasicResponse{Message: "Activity log was not found in provided activity"})
		return
	}

	activityLog := ac.Service.GetById(logId)
	c.JSON(200, services.ActivityLogModelToDto(activityLog))
}

func (ac *ActivityLogApi) GetAllActivityLogs(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityId := c.Param("activityId")
	if activityId == "" {
		c.JSON(400, data.BasicResponse{Message: "Activity Log Id is required"})
		return
	}
	if !ac.ActivityService.DoesUserOwnActivity(*userId, activityId) {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityLogs := ac.Service.GetAll(activityId)
	c.JSON(200, services.ActivityLogModelsToDtos(activityLogs))
}

func (ac *ActivityLogApi) CreateActivityLog(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityId := c.Param("activityId")
	if activityId == "" {
		c.JSON(400, data.BasicResponse{Message: "Activity Log Id is required"})
		return
	}
	if !ac.ActivityService.DoesUserOwnActivity(*userId, activityId) {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	var activityLogDto data.ActivityLogInDto
	if err := c.ShouldBindJSON(&activityLogDto); err != nil {
		c.JSON(400, data.BasicResponse{Message: "Malformed request body. " + err.Error()})
		return
	}
	logDate, timeParseError := time.Parse(services.DateTimeFormat, activityLogDto.LogDate)
	if timeParseError != nil {
		c.JSON(400, data.BasicResponse{Message: "Time format is not correct"})
	}
	activityLog := data.ActivityLog{
		Value:   activityLogDto.Value,
		LogDate: logDate,
	}
	createdLog := ac.Service.Create(activityId, activityLog)
	c.JSON(200, services.ActivityLogModelToDto(createdLog))
}

func (ac *ActivityLogApi) UpdateActivityLog(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityId, logId := c.Param("activityId"), c.Param("logId")
	if activityId == "" || logId == "" {
		c.JSON(400, data.BasicResponse{Message: "Malformed request"})
		return
	}
	if !ac.ActivityService.DoesUserOwnActivity(*userId, activityId) {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	if !ac.ActivityService.DoesActivityOwnsLog(activityId, logId, ac.Service) {
		c.JSON(404, data.BasicResponse{Message: "Activity log was not found in provided activity"})
		return
	}
	var activityLogDto data.ActivityLogInDto
	if err := c.ShouldBindJSON(&activityLogDto); err != nil {
		c.JSON(400, data.BasicResponse{Message: "Malformed request body. " + err.Error()})
		return
	}
	update, updateError := ac.Service.Update(logId, activityLogDto)
	if updateError != nil {
		c.JSON(400, "Failed to update activity"+updateError.Error())
		return
	}
	c.JSON(200, services.ActivityLogModelToDto(update))
}

func (ac *ActivityLogApi) DeleteActivityLog(c *gin.Context) {
	userId, authError := services.Authorize(c)
	if authError != nil {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	activityId, logId := c.Param("activityId"), c.Param("logId")
	if activityId == "" || logId == "" {
		c.JSON(400, data.BasicResponse{Message: "Malformed request"})
		return
	}
	if !ac.ActivityService.DoesUserOwnActivity(*userId, activityId) {
		c.JSON(401, data.BasicResponse{Message: "Unauthorized"})
		return
	}
	if !ac.ActivityService.DoesActivityOwnsLog(activityId, logId, ac.Service) {
		c.JSON(404, data.BasicResponse{Message: "Activity log was not found in provided activity"})
		return
	}
	c.JSON(200, ac.Service.Delete(logId))
}
