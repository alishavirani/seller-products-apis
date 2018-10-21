package main

import (
	"seller-backend/apis"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
)

func CreateRouter(session *mgo.Session) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", apis.GetAllProducts).Methods("GET")

	return router
}
