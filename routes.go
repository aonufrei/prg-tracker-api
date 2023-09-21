package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"prg-tracker/api"
	"prg-tracker/data"
	"prg-tracker/services"
)

type Dependencies struct {
	UserService        services.UserService
	ActivityService    services.ActivityService
	ActivityLogService services.ActivityLogService
	UserApi            api.UserApi
	ActivityApi        api.ActivityApi
	ActivityLogApi     api.ActivityLogApi
	LoginApi           api.LoginApi
}

func InitDependencies(db *gorm.DB) (*Dependencies, error) {
	userService := services.UserService{DB: db}
	activityService := services.ActivityService{DB: db}
	activityLogService := services.ActivityLogService{DB: db}
	userApi := api.UserApi{Service: &userService}
	activityApi := api.ActivityApi{Service: &activityService}
	activityLogApi := api.ActivityLogApi{Service: &activityLogService, ActivityService: &activityService}
	loginApi := api.LoginApi{UserService: &userService}
	return &Dependencies{
		UserService:        userService,
		ActivityService:    activityService,
		ActivityLogService: activityLogService,
		UserApi:            userApi,
		ActivityApi:        activityApi,
		ActivityLogApi:     activityLogApi,
		LoginApi:           loginApi,
	}, nil
}

func InitRoutes(root *gin.Engine, dependencies *Dependencies) {
	loginApi := dependencies.LoginApi
	userApi := dependencies.UserApi
	activityApi := dependencies.ActivityApi
	activityLogApi := dependencies.ActivityLogApi
	apiGroup := root.Group("api/v1/")

	apiGroup.POST("/login", loginApi.Login)

	requireAdmin := requireSpecificRole(data.ADMIN, dependencies.UserService)
	requireRegular := requireSpecificRole(data.REGULAR, dependencies.UserService)

	apiGroup.GET("/users/:id", requireAdmin, userApi.GetUser)
	apiGroup.GET("/users", requireAdmin, userApi.GetAllUsers)
	apiGroup.POST("/users", requireAdmin, userApi.CreateUser)
	apiGroup.PUT("/users/:id", requireAdmin, userApi.UpdateUser)
	apiGroup.DELETE("/users/:id", requireAdmin, userApi.DeleteUser)

	apiGroup.GET("/activities/:activityId", requireRegular, activityApi.GetActivity)
	apiGroup.GET("/activities", requireRegular, activityApi.GetAllActivities)
	apiGroup.POST("/activities", requireRegular, activityApi.CreateActivity)
	apiGroup.PUT("/activities/:activityId", requireRegular, activityApi.UpdateActivity)
	apiGroup.DELETE("/activities/:activityId", requireRegular, activityApi.DeleteActivity)

	apiGroup.GET("/activities/:activityId/log/:logId", requireRegular, activityLogApi.GetActivityLog)
	apiGroup.GET("/activities/:activityId/log", requireRegular, activityLogApi.GetAllActivityLogs)
	apiGroup.POST("/activities/:activityId/log", requireRegular, activityLogApi.CreateActivityLog)
	apiGroup.PUT("/activities/:activityId/log/:logId", requireRegular, activityLogApi.UpdateActivityLog)
	apiGroup.DELETE("/activities/:activityId/log/:logId", requireRegular, activityLogApi.DeleteActivityLog)

}
