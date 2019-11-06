package domain

import (
	"fmt"
	"net/http"

	"github.com/markgardner10/golang-microservice/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "iam", LastName: "testing", Email: "dummy@email.com"},
	}
)

func GetUser(userID int64) (*User, *utils.ApiError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApiError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
