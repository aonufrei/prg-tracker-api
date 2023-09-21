package data

import (
	"gorm.io/gorm"
	"time"
)

const (
	ADMIN   = "ADMIN"
	REGULAR = "REGULAR"
)

const (
	OCCURRENCE = "OCCURRENCE"
	COUNT      = "COUNT"
	MINUTES    = "MINUTES"
)

type User struct {
	gorm.Model
	Id       string `gorm:"primary_key"`
	Name     string
	Username string
	Password string
	Role     string
}

type Activity struct {
	gorm.Model
	Id        string `gorm:"primary_key"`
	Name      string
	UserId    string
	Type      string
	CreatedAt time.Time
}

type ActivityLog struct {
	gorm.Model
	Id         string `gorm:"primary_key"`
	Value      float32
	ActivityId string
	LogDate    time.Time
}
