package api

import(
	"net/http"
	"testing"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"bytes"
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
	
)

/*type Repository struct{}*/
/*type Controller struct{
	Repository api.Repository
}
var controller = &Controller{Repository: api.Repository{}}*/

//type Controller = Controller
//var controller = &Controller{Repository: Repository{}}

func sendRequest(request *http.Request,router *mux.Router) *httptest.ResponseRecorder{
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response

}
func TestWelcome(t *testing.T){
	request,_ := http.NewRequest("GET","/",nil)
	router := mux.NewRouter()
	router.Name("Welcome").Methods("GET").Path("/").HandlerFunc(controller.Welcome)
	response := sendRequest(request,router)
	assert.Equal(t,200,response.Code,"OK response is expected")
	var b string
	json.Unmarshal(response.Body.Bytes(), &b)
	assert.Equal(t,"Welcome to E-Commerce",b)


}
func TestViewAllProduct(t *testing.T){
	
	
	request,_ := http.NewRequest("GET","/Product",nil)
	router := mux.NewRouter()
	router.Name("ViewAllProduct").Methods("GET").Path("/Product").HandlerFunc(controller.ViewAllProduct)
	response := sendRequest(request,router)
	assert.Equal(t,200,response.Code,"OK response is expected")
	


}


func TestAddViewUpdateDeleteProduct(t *testing.T){

	//Test AddProduct Method
	var jstring = []byte(`{ "title": "Dummy","image": "dummy url","price": 1000, "rating": 0.0 }`)
	request,_ := http.NewRequest("POST","/Product/Add",bytes.NewBuffer(jstring))
	router := mux.NewRouter()
	router.Name("AddProduct").Methods("POST").Path("/Product/Add").HandlerFunc(controller.AddProduct)
	response := sendRequest(request,router)
	assert.Equal(t,200,response.Code,"OK response is expected")


	//Test ViewProduct
	var result Product
	session ,err := mgo.Dial(SERVER)
	if(err != nil){
		log.Fatalln("Database connection could not be established")
	}
	defer session.Close()
	
	session.DB(DBNAME).C(COLLECTION).Find(nil).Sort("-_id").Limit(1).One(&result)
	
	temp := "/Product/%d"
	temp = fmt.Sprintf(temp,result.Id)
	request,_ = http.NewRequest("GET",temp,nil)
	router = mux.NewRouter()
	router.Name("ViewProduct").Methods("GET").Path("/Product/{id}").HandlerFunc(controller.ViewProduct)
	response = sendRequest(request,router)
	assert.Equal(t,200,response.Code,"OK response is expected")
	temp = string(`{ "Id" : %d, "title" : %s, "image" : %s , "price" : %d, "rating" : %f}`)
	temp = fmt.Sprintf(temp,result.Id,string("\""+result.Title+"\""),string("\""+result.Image+"\""),result.Price,result.Rating)
	expected := (temp)
	assert.JSONEq(t, expected, response.Body.String(), "Response body differs")


	//Test UpdateProduct Method
	temp = `{ "Id" : %d , "title": "DummyGame","image": "dummy url","price": 1200, "rating": 4.0 }`
	temp = fmt.Sprintf(temp,result.Id)
    jstring = []byte(temp)
	request,_ = http.NewRequest("PUT","/Product/Update",bytes.NewBuffer(jstring))
	router = mux.NewRouter()
	router.Name("UpdateProduct").Methods("PUT").Path("/Product/Update").HandlerFunc(controller.UpdateProduct)
	response = sendRequest(request,router)
	assert.Equal(t,200,response.Code,"OK response is expected")


	temp = "/Product/Delete/%d"
	temp = fmt.Sprintf(temp,result.Id)
	request,_ = http.NewRequest("DELETE",temp,nil)
	router = mux.NewRouter()
	router.Name("DeleteProduct").Methods("DELETE").Path("/Product/Delete/{id}").HandlerFunc(controller.DeleteProduct)
	response = sendRequest(request,router)
	assert.Equal(t,200,response.Code,"OK response is expected")



}

