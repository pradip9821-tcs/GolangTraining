package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", signup).Methods("GET")

	log.Println("Server Listening on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("signup invoked.")
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 10

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 10
	vegetables["Beets"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
