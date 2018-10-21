package controllers

import (
	"seller-backend/structs"
	"seller-backend/utils"
)

func GetAllProducts() (structs.Products, error) {
	session := utils.ConnectToMongo()
	db := utils.LoadConfig().MongoDB.Db
	collection := utils.LoadConfig().MongoDB.Collections.Products

	conn := session.DB(db).C(collection)
	products := structs.Products{}

	err := conn.Find(nil).All(&products)

	return products, err
}
