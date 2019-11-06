package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/markgardner10/golang-microservice/mvc/services"
	"github.com/markgardner10/golang-microservice/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		apiErr := &utils.ApiError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonData, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)

		resp.Write(jsonData)
		return
	}
	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonData, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		// resp.Header().Set("Content-Type", "application/json")
		resp.Write(jsonData)
		return
	}
	jsonData, _ := json.Marshal(user)
	// resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}
