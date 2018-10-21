package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"seller-backend/controllers"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	//Try to use context

	products, err := controllers.GetAllProducts()
	if err != nil {
		//Write err on FE client
		log.Fatal("Error in fetching all products", err)
	}
	data, err := json.Marshal(products)
	if err != nil {
		//Write err on FE client
		log.Fatal("Error in marshalling all products", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
