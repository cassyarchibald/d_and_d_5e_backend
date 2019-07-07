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
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllClasses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show all classes")
	json.NewEncoder(w).Encode(Classes)
}

func main() {
	Classes = []Class{
		Class{Name: "Cassy"},
		Class{Name: "Cindy"},
	}
	handleRequests()
}
