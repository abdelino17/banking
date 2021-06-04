package app

import (
	"encoding/json"
	"fmt"
	"github.com/abdelino17/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

//type Customer struct {
//	Name    string `json:"full_name" xml:"name"`
//	City    string `json:"city" xml:"city"`
//	Zipcode string `json:"zip_code" xml:"zipcode"`
//}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Abdel", "Lome", "171009"},
	//	{"Aziz", "Evry", "91000"},
	//}

	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")

	//if r.Header.Get("Content-Type") == "application/xml" {
	//	w.Header().Add("Content-Type", "application/xml")
	//	_ = xml.NewEncoder(w).Encode(customers)
	//} else {
	//	w.Header().Add("Content-Type", "application/json")
	//	_ = json.NewEncoder(w).Encode(customers)
	//}
}