package api

import(
	"encoding/json"
	 "net/http"
	 "fmt"
	 "github.com/gorilla/mux"
	 "io/ioutil"
	 "strconv"
	 "strings"


)

type Controller struct{
	Repository Repository
}

//Welcome GET /
func (c *Controller) Welcome(w http.ResponseWriter, r *http.Request){

	
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	body,_ := json.Marshal("Welcome to E-Commerce")
	w.Write(body)
	return 
}

//ViewAllProduct GET /Product
func (c *Controller) ViewAllProduct(w http.ResponseWriter, r *http.Request){
	//Fetch all the products from Database

	products := c.Repository.GetAllProducts()
	if(products == (Products{}))
	{
		data,_ := json.Marshal("Failed to fetch the product details")

		//Set the response header 
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		w.WriteHeader(StatusInternalServerError)

		//Write the product details
		w.Write(data)
		return
	}
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
		fmt.Println("Error reading the request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &product); err != nil{
		fmt.Println("Error reading the request body")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	data,_ := json.Marshal("success")
	w.Write(data)
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
		fmt.Println("Error while processing request ")
		return
	}

	//Get the product information for the input id
	product := c.Repository.GetProduct(product_id)
	if(product == (Products{}))
	{

	}

	//Convert the structure to json object
	data,_ := json.Marshal(product)

	//Set the http header and status
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return

}

//UpdateProduct PUT /Product/Update
func (c *Controller) UpdateProduct(w http.ResponseWriter,r *http.Request){
	var product Product
	//Read the request body
	body,err := ioutil.ReadAll(r.Body)

	if err!= nil{
		fmt.Println("Error reading the request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &product); err != nil{
		fmt.Println("Error reading the request body")
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	success := c.Repository.UpdateProduct(product)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
   data,_ := json.Marshal("success")
	w.Write(data)
    return


}

func (c *Controller) DeleteProduct(w http.ResponseWriter,r *http.Request){
	//Get all the input variables
	vars := mux.Vars(r)

	//Get the product id from the variables
	id := vars["id"]

	//Convert id to integer
	product_id,err := strconv.Atoi(id);

	if err != nil{
		fmt.Println("Error while processing request ")
	}

	if err := c.Repository.DeleteProduct(product_id); err != "OK" { // delete a product by id
        fmt.Println(err);
        if strings.Contains(err, "404") {
            w.WriteHeader(http.StatusNotFound)
        } else if strings.Contains(err, "500") {
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    data,_ := json.Marshal("success")
	w.Write(data)
    return

}




