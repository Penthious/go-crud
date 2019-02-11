package user

import (
	"go-crud/api/product_api"
	"go-crud/api/user_api"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/api/product/findall", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/user/create", user_api.Create).Methods("POST")
	router.HandleFunc("/api/user", user_api.Index).Methods("GET")
}
