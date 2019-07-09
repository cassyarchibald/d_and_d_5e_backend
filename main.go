package main

import (
	class "dndclass"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Classes
var Classes []class.DndClass

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

func loadClasses() {
	response, err := http.Get("https://api-beta.open5e.com/classes/")
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
		return
	} else {
		// var result []DndClass
		// data[results].each do |result|
		// 	result << DndClass.new(resultl[name])
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func returnAllClasses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show all classes")
	loadClasses()
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
	Classes = []DndClass{
		DndClass{Name: "Cassy"},
		DndClass{Name: "Cindy"},
	}
	handleRequests()
}
