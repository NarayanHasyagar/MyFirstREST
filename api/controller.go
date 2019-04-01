package api

import(
	"encoding/json"
	 "net/http"
	 "fmt"
	 "github.com/gorilla/mux"
	 "io/ioutil"
	 "log"
	 "strconv"


)

type Controller struct{
	Repository Repository
}

//Welcome GET /
func (c *Controller) Welcome(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"Welcome to E-Commerce")
	return
}

//ViewAllProduct GET /Product
func (c *Controller) ViewAllProduct(w http.ResponseWriter, r *http.Request){
	//Fetch all the products from Database

	products := c.Repository.GetAllProducts()
	//Convert the structure to json object
	data,_ := json.Marshal(products)

	//Set the response header 
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//Write the product details
	w.Write(data)
	return
}

//AddProduct POST /Product/Add
func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request){

	var product Product
	//Read the request body
	body,err := ioutil.ReadAll(r.Body)


	if err!= nil{
		log.Fatalln("Error reading the request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &product); err != nil{
		log.Fatalln("Error reading the request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Add the product to database
	success := c.Repository.AddProduct(product)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Set the response header
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return

}

//ViewProduct GET /Product/{id}
func (c *Controller) ViewProduct(w http.ResponseWriter,r *http.Request){
	//Get all the input variables
	vars := mux.Vars(r)

	//Get the product id from the variables
	id := vars["id"]

	//Convert id to integer
	product_id,err := strconv.Atoi(id);

	if err != nil{
		log.Fatalln("Error while processing request ")
	}

	//Get the product information for the input id
	product := c.Repository.GetProduct(product_id)

	//Convert the structure to json object
	data,_ := json.Marshal(product)

	//Set the http header and status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return

}




