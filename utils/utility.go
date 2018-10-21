package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"seller-backend/structs"

	"github.com/globalsign/mgo"
)

// type session *mgo.Session

// var sess session

// // GetMyKey returns a value for this package from the request values.
// func GetMyKey(r *http.Request) session {
// 	if rv := context.Get(r, sess); rv != nil {
// 		fmt.Println("RV??", rv)
// 		return rv.(session)
// 	}
// 	return nil
// }

// // SetMyKey sets a value for this package in the request values.
// func SetMyKey(r *http.Request, val *mgo.Session) {
// 	context.Set(r, sess, val)
// }

func ConnectToMongo() *mgo.Session {
	fmt.Println("enter main - connecting to mongo")

	// maxWait := time.Duration(5 * time.Second)
	session, sessionErr := mgo.Dial("localhost:27017")

	if sessionErr != nil {
		log.Fatal("Error in connecting to MongoDb", sessionErr)
	}
	return session
}

func LoadConfig() structs.Config {
	config := structs.Config{}

	file, err := os.Open("../config/config.development.json")
	if err != nil {
		log.Fatal("Error in opening config file ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error in loading config:", err)
	}
	return config
}
