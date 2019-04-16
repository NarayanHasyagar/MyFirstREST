package api

import(
	"log"
	"gopkg.in/mgo.v2"
	"fmt"
	
)

type Repository struct{}


//Localhost server url
const SERVER = "mongodb://host.docker.internal:27017/Ecomm"

const DBNAME = "Ecomm"

const COLLECTION = "product"

var productId = 22;

//Fetch all the products from the database
func (r Repository) GetAllProducts() Products{
	//Establish connection with mongo database
	session , err := mgo.Dial(SERVER)
	if(err != nil){
		log.Fatalln("Database connection could not be established")
	}
	defer session.Close()

	//Get the session object for the required and database collection
	c := session.DB(DBNAME).C(COLLECTION)

	//Refer the model structure template
	resultset := Products{}

	//Execute the mongo query for getting the results
	if err := c.Find(nil).All(&resultset); err != nil{
			log.Fatalln("Failed to get the results")
	}
	return resultset
}

//Fetch details for a perticular product
func (r Repository) GetProduct(id int)Product{
	session,err := mgo.Dial(SERVER)

	if(err != nil){
		log.Fatalln("Database connection could not be established")
	}
	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)

	var result Product

	if err:= c.FindId(id).One(&result); err != nil{
		log.Fatalln("Failed to get the product details")
	}

	return result
}

//Add product details to database
func (r Repository) AddProduct(product Product) bool{

	session,err := mgo.Dial(SERVER)

	if(err != nil){
		log.Fatalln("Database connection could not be established")
	}
	defer session.Close()

	//Increase product Id
	productId = productId + 1
	product.Id = productId

	if err := session.DB(DBNAME).C(COLLECTION).Insert(product); err != nil{
		log.Fatalln(err)
		return false
	}
	fmt.Println("Success")
	return true

}

//Update product details in database
func (r Repository) UpdateProduct(product Product) bool{
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	err = session.DB(DBNAME).C(COLLECTION).UpdateId(product.Id, product)
	

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Success")
	return true
}

func (r Repository) DeleteProduct(id int) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Remove product
	if err = session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Success")
	// Write status
	return "OK"
}

