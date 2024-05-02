package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/MaxRubel/kennel-server-go/views"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/animals", func(w http.ResponseWriter, r *http.Request) {
		urlString := r.URL.String()
		parsedURL, err := url.Parse(urlString)
		fmt.Print(parsedURL)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			return
		}
		animalsJson, err := views.GetAllAnimals()
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(animalsJson)
	})

	r.HandleFunc("/animals/{id}", func(w http.ResponseWriter, r *http.Request) {
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
	})
	views.TestRequest()
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
