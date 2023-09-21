package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"prg-tracker/api"
	"prg-tracker/services"
)

type Dependencies struct {
	UserService        services.UserService
	ActivityService    services.ActivityService
	ActivityLogService services.ActivityLogService
	UserApi            api.UserApi
	ActivityApi        api.ActivityApi
	ActivityLogApi     api.ActivityLogApi
}

func InitDependencies(db *gorm.DB) (*Dependencies, error) {
	userService := services.UserService{DB: db}
	activityService := services.ActivityService{DB: db}
	activityLogService := services.ActivityLogService{DB: db}
	userApi := api.UserApi{Service: &userService}
	activityApi := api.ActivityApi{Service: &activityService}
	activityLogApi := api.ActivityLogApi{Service: &activityLogService, ActivityService: &activityService}
	return &Dependencies{
		UserService:        userService,
		ActivityService:    activityService,
		ActivityLogService: activityLogService,
		UserApi:            userApi,
		ActivityApi:        activityApi,
		ActivityLogApi:     activityLogApi,
	}, nil
}

func InitRoutes(root *gin.Engine, dependencies *Dependencies) {
	userApi := dependencies.UserApi
	activityApi := dependencies.ActivityApi
	activityLogApi := dependencies.ActivityLogApi
	apiGroup := root.Group("api/v1/")
	apiGroup.GET("users/:id", userApi.GetUser)
	apiGroup.GET("/users", userApi.GetAllUsers)
	apiGroup.POST("/users", userApi.CreateUser)
	apiGroup.PUT("/users/:id", userApi.UpdateUser)
	apiGroup.DELETE("/users/:id", userApi.DeleteUser)

	apiGroup.GET("activities/:activityId", activityApi.GetActivity)
	apiGroup.GET("/activities", activityApi.GetAllActivities)
	apiGroup.POST("/activities", activityApi.CreateActivity)
	apiGroup.PUT("/activities/:activityId", activityApi.UpdateActivity)
	apiGroup.DELETE("/activities/:activityId", activityApi.DeleteActivity)

	apiGroup.GET("activities/:activityId/log/:logId", activityLogApi.GetActivityLog)
	apiGroup.GET("/activities/:activityId/log", activityLogApi.GetAllActivityLogs)
	apiGroup.POST("/activities/:activityId/log", activityLogApi.CreateActivityLog)
	apiGroup.PUT("/activities/:activityId/log/:logId", activityLogApi.UpdateActivityLog)
	apiGroup.DELETE("/activities/:activityId/log/:logId", activityLogApi.DeleteActivityLog)
}
