package data

import "time"

type BasicResponse struct {
	Message string `json:"message"`
}

type UserOutDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserInDto struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ActivityOutDto struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	UserId    string    `json:"user_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type ActivityInDto struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ActivityLogOutDto struct {
	Id         string    `json:"id"`
	Value      float32   `json:"value"`
	ActivityId string    `json:"activity_id"`
	LogDate    time.Time `json:"log_date"`
}

type ActivityLogInDto struct {
	Value   float32 `json:"value"`
	LogDate string  `json:"log_date"`
}
