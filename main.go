package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MaxRubel/kennel-server-go/handle"
	"github.com/MaxRubel/kennel-server-go/models"
	"github.com/MaxRubel/kennel-server-go/views"
	"github.com/gorilla/mux"
)

func writeCORS(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	r := mux.NewRouter()

	// ANIMAL REQUESTS
	r.HandleFunc("/animals", func(w http.ResponseWriter, r *http.Request) {
	
		writeCORS(w)

		if r.Method == "GET" {
			animalsJson, err := views.GetAllAnimals()
			if err != nil {
				panic(err)
			}
			w.Write(animalsJson)
		}
		if r.Method == "POST" {
			var data models.Animal
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			err = views.CreateNewAnimal(data)
			if err != nil {
				fmt.Print("Error creating new animal at the DB")
				return
			}
			w.WriteHeader(http.StatusCreated)
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
		writeCORS(w)
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
	handle.TestErrorDb()
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
