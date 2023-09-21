package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"prg-tracker/data"
	"time"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
)

type ActivityLogService struct {
	DB *gorm.DB
}

func (as *ActivityLogService) GetById(id string) *data.ActivityLog {
	activityLog := data.ActivityLog{}
	result := as.DB.Where(&data.ActivityLog{Id: id}).First(&activityLog)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &activityLog
}

func (as *ActivityLogService) GetAll(activityId string) []data.ActivityLog {
	activityLogs := make([]data.ActivityLog, 0)
	as.DB.Where(&data.ActivityLog{ActivityId: activityId}).Find(&activityLogs)
	return activityLogs
}

func (as *ActivityLogService) Create(activityId string, activityLog data.ActivityLog) *data.ActivityLog {
	activityLog.Id = uuid.NewString()
	activityLog.ActivityId = activityId
	as.DB.Create(&activityLog)
	return &activityLog
}

func (as *ActivityLogService) Update(id string, inActivity data.ActivityLogInDto) (*data.ActivityLog, error) {
	activityById := as.GetById(id)
	if activityById == nil {
		return nil, errors.New("activity_log was not found")
	}
	logDate, timeParseError := time.Parse(DateTimeFormat, inActivity.LogDate)
	if timeParseError != nil {
		return nil, timeParseError
	}
	activityById.Value = inActivity.Value
	activityById.LogDate = logDate
	as.DB.Save(activityById)
	return activityById, nil
}

func (as *ActivityLogService) Delete(id string) bool {
	activityLog := as.GetById(id)
	if activityLog == nil {
		return false
	}
	as.DB.Delete(&activityLog)
	return true
}

func (as *ActivityLogService) DoUserOwnLog(userId, logId string, activityService *ActivityService) bool {
	log := as.GetById(logId)
	if log == nil {
		return false
	}
	activity := activityService.GetById(log.ActivityId)
	if activity == nil {
		return false
	}
	return activity.UserId == userId
}
