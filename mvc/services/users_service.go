package services

import (
	"github.com/markgardner10/golang-microservices/mvc/domain"
	"github.com/markgardner10/golang-microservices/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApiError) {
	return domain.GetUser(userID)
}
