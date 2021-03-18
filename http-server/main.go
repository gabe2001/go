package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Contact struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

var contacts []Contact

var r = mux.NewRouter()
var server http.Server

func main() {
	initData()
	r.HandleFunc("/contacts", getContacts).Methods("GET")
	r.HandleFunc("/contacts/{name}", getContact).Methods("GET")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	r.HandleFunc("/contacts/{name}", updateContact).Methods("PUT")
	r.HandleFunc("/contacts/{name}", deleteContact).Methods("DELETE")
	r.HandleFunc("/shutdown", shutdown).Methods("GET")
	startServer()
	fmt.Println("exit")
}

func initData() {
	contacts = append(contacts, Contact{Name: "Antoine", Age: 52})
	contacts = append(contacts, Contact{Name: "Christoph", Age: 50})
	contacts = append(contacts, Contact{Name: "Gabriel", Age: 56})
	contacts = append(contacts, Contact{Name: "Walter", Age: 58})
}

// https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve
func startServer() {
	server := &http.Server{Addr: ":8085", Handler: r}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func shutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
	shutdownServer()
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		fmt.Println(idx)
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	json.NewDecoder(r.Body).Decode(&contact)
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		fmt.Println(idx)
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			var contact Contact
			json.NewDecoder(r.Body).Decode(&contact)
			contact.Name = params["name"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}

}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		fmt.Println(idx)
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}
