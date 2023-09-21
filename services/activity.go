package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"prg-tracker/data"
)

type ActivityService struct {
	DB *gorm.DB
}

func (as *ActivityService) GetById(id string) *data.Activity {
	activity := data.Activity{}
	result := as.DB.Where(&data.Activity{Id: id}).First(&activity)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &activity
}

func (as *ActivityService) GetByUserIdAndName(userId, name string) *data.Activity {
	activity := data.Activity{}
	result := as.DB.Where(&data.Activity{UserId: userId, Name: name}).First(&activity)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &activity
}

func (as *ActivityService) GetAll(userId string) []data.Activity {
	activities := make([]data.Activity, 0)
	as.DB.Where(&data.Activity{UserId: userId}).Find(&activities)
	return activities
}

func (as *ActivityService) Create(userId string, activity data.Activity) *data.Activity {
	activity.Id = uuid.NewString()
	activity.UserId = userId
	as.DB.Create(&activity)
	return &activity
}

func (as *ActivityService) Update(id string, inActivity data.ActivityInDto) (*data.Activity, error) {
	activityById := as.GetById(id)
	if activityById == nil {
		return nil, errors.New("activity was not found")
	}
	activityById.Name = inActivity.Name
	activityById.Type = inActivity.Type
	as.DB.Save(activityById)
	return activityById, nil
}

func (as *ActivityService) Delete(id string) bool {
	activity := as.GetById(id)
	if activity == nil {
		return false
	}
	as.DB.Delete(&activity)
	return true
}

func (as *ActivityService) DoesActivityOwnsLog(activityId, logId string, activityLogService *ActivityLogService) bool {
	activity := as.GetById(activityId)
	if activity == nil {
		return false
	}
	log := activityLogService.GetById(logId)
	if log == nil {
		return false
	}
	return log.ActivityId == activityId
}

func (as *ActivityService) DoesUserOwnActivity(userId, activityId string) bool {
	activity := as.GetById(activityId)
	if activity == nil {
		return false
	}
	return activity.UserId == userId
}

func (as *ActivityService) IsNameUnique(userId, name string) bool {
	activity := as.GetByUserIdAndName(userId, name)
	return activity == nil
}

func (as *ActivityService) IsNameUniqueForUpdate(activityId, userId, name string) bool {
	activity := as.GetByUserIdAndName(userId, name)
	if activity == nil {
		return true
	} else {
		return activity.Id == activityId
	}
}
