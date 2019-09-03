package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Partner struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Partners array
// to simulate a database
var Partners []Partner

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to partner service!")
	fmt.Println("Endpoint: homePage")
}

func createNewPartner(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var partner Partner
	// unmarshall into a new partner struct
	json.Unmarshal(reqBody, &partner)
	Partners = append(Partners, partner)
	json.NewEncoder(w).Encode(partner)
}

func returnAllPartners(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPartners")
	json.NewEncoder(w).Encode(Partners)
}

func returnSinglePartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, partner := range Partners {
		if partner.ID == key {
			json.NewEncoder(w).Encode(partner)
		}
	}
}

func deletePartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, partner := range Partners {
		if partner.ID == id {
			Partners = append(Partners[:index], Partners[index+1:]...)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/partner", createNewPartner).Methods("POST")
	router.HandleFunc("/partners", returnAllPartners)
	router.HandleFunc("/partner/{id}", returnSinglePartner)
	router.HandleFunc("/partner/{id}", deletePartner).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	fmt.Println("Rest API v0.1 - Mux Routers")
	Partners = []Partner{
		Partner{ID: "1", Title: "Antony", Desc: "This is a ", Content: "Partner Content"},
		Partner{ID: "2", Title: "Mr T", Desc: "This is another test", Content: "Content guranteed"},
	}
	handleRequests()
}
