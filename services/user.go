package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"prg-tracker/data"
)

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) GetById(id string) *data.User {
	user := data.User{}
	result := us.DB.Where("Id = ?", id).First(&user)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &user
}

func (us *UserService) GetByUsername(username string) *data.User {
	var user data.User
	result := us.DB.Where("username = ?", username).First(&user)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &user
}

func (us *UserService) GetAll() []data.User {
	users := make([]data.User, 0)
	us.DB.Find(&users)
	return users
}

func (us *UserService) Create(user data.User) *data.User {
	user.Id = uuid.NewString()
	us.DB.Create(&user)
	return &user
}

func (us *UserService) Update(id string, inUser data.UserInDto) (*data.User, error) {
	userById := us.GetById(id)
	if userById == nil {
		return nil, errors.New("user was not found")
	}
	userById.Name = inUser.Name
	userById.Username = inUser.Username
	userById.Password = inUser.Password
	userById.Role = inUser.Role
	us.DB.Save(userById)
	return userById, nil
}

func (us *UserService) Delete(id string) bool {
	user := us.GetById(id)
	if user == nil {
		return false
	}
	us.DB.Delete(&user)
	return true
}
