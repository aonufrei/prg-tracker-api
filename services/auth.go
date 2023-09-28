package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"prg-tracker/data"
	"strings"
	"time"
)

const (
	AuthHeader = "Authorization"

	DefaultTokenLifespan = 2 * 24 * time.Hour // 2 days
)

func ProcessLoginRequest(request data.LoginInDto, service *UserService) (string, error) {
	if request.Username == nil && request.Password == nil {
		return "", errors.New("both username and password is required")
	}
	providedUsername := *request.Username
	providedPassword := *request.Password
	user := service.GetByUsername(providedUsername)
	if user == nil {
		return "", errors.New("user not found")
	}
	authorized := user.Password == providedPassword
	if !authorized {
		return "", errors.New("unauthorized")
	}
	authData := CreateAuthData(user.Id, DefaultTokenLifespan)
	return CreateToken(authData)
}

func Authorize(c *gin.Context) (*string, error) {
	token := c.GetHeader(AuthHeader)
	if token == "" {
		return nil, errors.New("no authorization token provided")
	}
	return DecodeAuthToken(token)
}

func CreateToken(authData data.AuthData) (string, error) {
	return encodeToBase64(authData)
}

func CreateAuthData(userId string, duration time.Duration) data.AuthData {
	expiration := time.Now().Add(duration)
	return data.AuthData{
		UserId:         userId,
		ExpirationDate: expiration,
	}
}

func DecodeAuthToken(token string) (*string, error) {
	parts := strings.Split(token, " ")
	if len(parts) != 2 {
		return nil, errors.New("bad authorization token provided")
	}
	authData := data.AuthData{}
	decodeError := decodeFromBase64(parts[1], &authData)
	if decodeError != nil || authData.ExpirationDate.Before(time.Now()) {
		return nil, errors.New("unauthorized")
	}
	return &authData.UserId, nil
}

func encodeToBase64[T interface{}](v T) (string, error) {
	jsonString, marshalErr := json.Marshal(v)
	if marshalErr != nil {
		return "", errors.New("failed to create token")
	}
	encodedStr := base64.StdEncoding.EncodeToString(jsonString)
	return encodedStr, nil
}

func decodeFromBase64[T interface{}](encodedStr string, v *T) error {
	jsonString, decodeErr := base64.StdEncoding.DecodeString(encodedStr)
	if decodeErr != nil {
		return errors.New("malformed auth token")
	}
	return json.Unmarshal(jsonString, v)
}
