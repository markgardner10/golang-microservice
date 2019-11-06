package domain

import (
	"fmt"
	"net/http"

	"github.com/markgardner10/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Mark", LastName: "Gardner", Email: "dummy@email.com"},
	}
)

func GetUser(userID int64) (*User, *utils.ApiError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApiError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
