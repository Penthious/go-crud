package product_api

import (
	"go-crud/config"
	"net/http"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"go-crud/models"
// )

func FindAll(response http.ResponseWriter, request *http.Request) {
	// 	fmt.Println("We are here")

	config.Connect()

	// 	if err != nil {
	// 		respondWithError(response, http.StatusBadRequest, err.Error())
	// 		return
	// 	}
	// 	productModel := models.ProductModel{
	// 		Db: db,
	// 	}

	// 	products, err2 := productModel.FindAll()

	// 	if err2 != nil {
	// 		respondWithError(response, http.StatusBadRequest, err2.Error())
	// 		return
	// 	}

	// 	respondWithJson(response, http.StatusOK, products)
}

// func respondWithError(w http.ResponseWriter, code int, msg string) {
// 	respondWithJson(w, code, map[string]string{"error": msg})
// }

// func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
