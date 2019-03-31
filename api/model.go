package api

type Product struct{
	Id string 'bson:"id"'
	Name string 'json:"name"'
	Title string 'json:"image"'
	Price uint64 'json:"price"'
	Rating uint8 'json:"rating"'
}

type Products []Product