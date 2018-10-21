package main

import (
	"fmt"
	"log"
	"net/http"

	"seller-backend/utils"
)

func main() {
	session := utils.ConnectToMongo()
	fmt.Println("Connected to MongoDB successfully!!!!")

	router := CreateRouter(session)

	//Can use TOML for loading config
	port := utils.LoadConfig().Server.Port

	//Add Support for CORS
	log.Fatal(http.ListenAndServe(port, router))
}
