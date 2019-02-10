package user_api

import (
	"encoding/json"
	"go-crud/config"
	"go-crud/entities"
	"net/http"
)

func Create(response http.ResponseWriter, request *http.Request) {
	// 	fmt.Println("We are here")

	db, err := config.Connect()
	defer db.Close()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	user := entities.User{Age: 32, Name: "teut", Email: "test@euost.com"}

	if err := db.Create(&user).Error; err != nil {
		respondWithJson(response, http.StatusBadRequest, err)
		return

	}

	respondWithJson(response, http.StatusOK, user)

}

func Index(response http.ResponseWriter, request *http.Request) {
	var users []entities.User
	db, err := config.Connect()
	defer db.Close()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	db.Find(&users)
	respondWithJson(response, http.StatusOK, users)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
