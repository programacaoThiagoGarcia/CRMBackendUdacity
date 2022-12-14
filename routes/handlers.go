package routes

import (
	"CRMBackend/Udacity/bd"
	"CRMBackend/Udacity/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var dataBase = bd.Database

// GetCustomers get all customers on database
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataBase)

}

// GetCustomer get one customer on database, that id is passed on URL request
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	item := mux.Vars(r)
	id, _ := strconv.ParseInt(item["id"], 10, 64)

	var customer model.Customer

	for _, item := range dataBase {
		if int64(item.ID) == id {
			customer = item
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// AddCustomer add one customer in database
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var customer model.Customer

	err = json.Unmarshal(body, &customer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lastId := dataBase[len(dataBase)-1].ID
	customer.ID = lastId + 1

	dataBase = append(dataBase, customer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dataBase)

}

// UpdateCustomer update customer information
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer model.Customer

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &customer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	for index, item := range dataBase {
		if item.ID == uint64(id) {
			customer.ID = uint64(id)
			dataBase[index] = customer
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dataBase)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item not Found"})

}

// DeleteCustomer delete customer information
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	for index, item := range dataBase {
		if item.ID == uint64(id) {
			dataBase = dataBase.Remove(index)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dataBase)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item not Found"})

}
