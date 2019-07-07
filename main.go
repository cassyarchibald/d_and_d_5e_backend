package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Class - Struct for D&D Classes
type Class struct {
	Name string `json:"name"`
}

var Classes []Class

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/classes", returnAllClasses)
	myRouter.HandleFunc("/classes/{name}", returnSingleClass)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllClasses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show all classes")
	json.NewEncoder(w).Encode(Classes)
}

func returnSingleClass(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show one class")

	vars := mux.Vars(r)
	key := vars["name"]

	for _, class := range Classes {
		if class.Name == key {
			json.NewEncoder(w).Encode(class)
		}
	}
}

func main() {
	Classes = []Class{
		Class{Name: "Cassy"},
		Class{Name: "Cindy"},
	}
	handleRequests()
}
