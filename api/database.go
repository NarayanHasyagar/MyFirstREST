package api

import(
	"log"
	"gopkg.in/mgo.v2"
	"fmt"
	
)

type Repository struct{}


//Localhost server url
//const SERVER = "mongodb://host.docker.internal:27058/sradb"
const SERVER = "mongodb://sraapp:chanGmen0w@64.101.4.199:27058/sradb"
const DBNAME = "sradb"

const COLLECTION = "product"

//var productId = 22;

//Fetch all the products from the database
func (r Repository) GetAllProducts() Products{
	//Refer the model structure template
	resultset := Products{}
	//Establish connection with mongo database
	session , err := mgo.Dial(SERVER)
	if(err != nil){
		fmt.Println("Database connection could not be established")
		return resultset
	}
	defer session.Close()

	//Get the session object for the required and database collection
	c := session.DB(DBNAME).C(COLLECTION)

	
	

	//Execute the mongo query for getting the results
	if err := c.Find(nil).All(&resultset); err != nil{
			fmt.Println("Failed to get the results")
	}
	return resultset
}

//Fetch details for a perticular product
func (r Repository) GetProduct(id int)Product{
	session,err := mgo.Dial(SERVER)
	defer session.Close()
	if(err != nil){
		log.Fatalln(err)
		log.Fatalln("Database connection could not be established")
	}
	

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
	defer session.Close()
	if(err != nil){
		log.Fatalln("Database connection could not be established")
	}
	

	//Increase product Id
	var result Product
	
	session.DB(DBNAME).C(COLLECTION).Find(nil).Sort("-_id").Limit(1).One(&result)
	
	result.Id = result.Id+1
	product.Id = result.Id

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
		fmt.Println(err)
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
		fmt.Println(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Success")
	// Write status
	return "OK"
}

