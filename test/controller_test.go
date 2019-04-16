package api

import(
	"net/http"
	"testing"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)


func sendRequest(request *http.Request,router *mux.Router) *httptest.ResponseRecorder{
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response

}
func TestWelcome(t *Testing.T)
{
	request,_ := http.NewRequest("GET","/",nil)
	router := mux.NewRouter()
	router.Name("Welcome").Methods("GET").Path("/").HandlerFunc(controller.Welcome)
	response := sendRequest(request,router)
	assert.Equal(t,200,response.code,"OK response is expected")
	var b string
	json.Unmarshal(response.Body.Bytes(), &b)
	assert.Equal(t,"Welcome to E-Commerce",b)


}