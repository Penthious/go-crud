package user_api

import (
	"encoding/json"
	"fmt"
	"go-crud/api/utils"
	"go-crud/config"
	"go-crud/entities"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var users []entities.User
	db, err := config.Connect()
	defer db.Close()

	if err != nil {
		utils.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	db.Find(&users)
	utils.RespondWithJson(response, http.StatusOK, users)
}

func Create(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		utils.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(request.Body)

	// user := entities.User{Age: 32, Name: "teut", Email: "test@euosue.com"}

	if err := db.Create(&user).Error; err != nil {
		utils.RespondWithJson(response, http.StatusBadRequest, err)
		return

	}

	utils.RespondWithJson(response, http.StatusOK, user)

}
