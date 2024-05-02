package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MaxRubel/kennel-server-go/views"
	"github.com/gorilla/mux"
)

func wildTest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println("Received request for:", path)
	fmt.Fprintf(w, "Hello, you requested: %s\n", path)
}

func main() {
	r := mux.NewRouter()

	// ANIMAL REQUESTS
	r.HandleFunc("/animals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			animalsJson, err := views.GetAllAnimals()
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(animalsJson)
		}
	})

	r.HandleFunc("/animals/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			vars := mux.Vars(r)
			idStr := vars["id"]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			animalJson, err := views.GetSingleAnimal(id)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(animalJson)
		}
	})

	//LOCATION REQUESTS
	r.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			locationsJson, err := views.GetAllLocations()
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(locationsJson)
		}
	})

	r.HandleFunc("/locations/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			vars := mux.Vars(r)
			idStr := vars["id"]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			locationJson, err := views.GetSingleLocation(id)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(locationJson)
		}
	})

	//CUSTOMER REQUESTS
	r.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			animalsJson, err := views.GetAllCustomers()
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(animalsJson)
		}
	})

	r.HandleFunc("/customers/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			vars := mux.Vars(r)
			idStr := vars["id"]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			customerJson, err := views.GetSingleCustomer(id)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(customerJson)
		}
	})

	//EMPLOYEE REQUESTS
	r.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			employeesJson, err := views.GetAllEmployees()
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(employeesJson)
		}
	})

	r.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			vars := mux.Vars(r)
			idStr := vars["id"]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			employeesJson, err := views.GetSingleCustomer(id)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(employeesJson)
		}
	})

	views.TestRequest()
	r.HandleFunc("/", wildTest)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
