package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"serge/food-lookup-demo/api"
)

const (
	productLimit = 100
)

func lookupFood(w http.ResponseWriter, r *http.Request) {

	q := r.FormValue("q")

	log.Printf("lookup like {%s} products requested\n", q)

	ps, err := api.GetProducts(q, productLimit)

	if err != nil {
		log.Fatal(err)
	}

	responseJSON, err := json.Marshal(ps)

	if err != nil {
		log.Printf("[ERROR] err = %+v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := string(responseJSON)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, response)

}

func main() {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/food", lookupFood).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
