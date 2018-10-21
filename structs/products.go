package structs

type Product struct {
	ID       int    `bson:"_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Products []Product
