package services

import (
	"errors"
	"prg-tracker/data"
)

func UserModelToDto(model *data.User) data.UserOutDto {
	return data.UserOutDto{
		Id:       model.Id,
		Name:     model.Name,
		Username: model.Username,
		Role:     model.Role,
	}
}

func UserModelsToDtos(models []data.User) []data.UserOutDto {
	dtos := make([]data.UserOutDto, len(models))
	for i, v := range models {
		dtos[i] = UserModelToDto(&v)
	}
	return dtos
}

func ActivityModelToDto(model *data.Activity) data.ActivityOutDto {
	return data.ActivityOutDto{
		Id:        model.Id,
		Name:      model.Name,
		UserId:    model.UserId,
		Type:      model.Type,
		CreatedAt: model.CreatedAt,
	}
}

func ActivityModelsToDtos(models []data.Activity) []data.ActivityOutDto {
	dtos := make([]data.ActivityOutDto, len(models))
	for i, v := range models {
		dtos[i] = ActivityModelToDto(&v)
	}
	return dtos
}

func ActivityLogModelToDto(model *data.ActivityLog) data.ActivityLogOutDto {
	return data.ActivityLogOutDto{
		Id:         model.Id,
		Value:      model.Value,
		ActivityId: model.ActivityId,
		LogDate:    model.LogDate,
	}
}

func ActivityLogModelsToDtos(models []data.ActivityLog) []data.ActivityLogOutDto {
	dtos := make([]data.ActivityLogOutDto, len(models))
	for i, v := range models {
		dtos[i] = ActivityLogModelToDto(&v)
	}
	return dtos
}

func ValidateUserInDto(dto data.UserInDto) error {
	nameLen := len(dto.Name)
	if nameLen < 3 || nameLen > 30 {
		return errors.New("user name is not in the valid format")
	}
	usernameLen := len(dto.Username)
	if usernameLen <= 5 || usernameLen > 40 {
		return errors.New("username is not in the valid format")
	}
	passwordLen := len(dto.Password)
	if passwordLen <= 8 || passwordLen > 40 {
		return errors.New("password is not in the valid format")
	}
	role := dto.Role
	if err := ValidateUserRole(role); err != nil {
		return err
	}
	return nil
}

func ValidateUsername(username string, service *UserService) error {
	user := service.GetByUsername(username)
	if user != nil {
		return errors.New("provided username is already in use")
	}
	return nil
}

func ValidateUserIdForUpdate(id string, service *UserService) error {
	if service.GetById(id) == nil {
		return errors.New("user with id does not exist")
	}
	return nil
}

func ValidateUsernameForUser(id, username string, service *UserService) error {
	userByUsername := service.GetByUsername(username)
	if userByUsername == nil {
		return nil
	}
	if userByUsername.Id != id {
		return errors.New("provided username is already in use")
	}
	return nil
}

func ValidateUserRole(userRole string) error {
	for _, t := range data.GetAllUserRoles() {
		if t != userRole {
			return errors.New("none existing User Role was provided")
		}
	}
	return nil
}

func ValidateActivityType(activityType string) error {
	for _, t := range data.GetAllActivityTypes() {
		if t != activityType {
			return errors.New("none existing Activity Type was provided")
		}
	}
	return nil
}
