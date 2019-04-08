package api

//Structure to handle the json data
type Product struct{
	Id int `bson:"_id,omitempty"`
	Title string `json:"title"`
	Image string `json:"image"`
	Price uint64 `json:"price"`
	Rating float32 `json:"rating"`
}

type Products []Product