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
	_, err := fmt.Println("exit")
	if err != nil {
		return
	}
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

func shutdown(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncode(w, "")
	shutdownServer()
}

func getContacts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncode(w, contacts)
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		_, err := fmt.Println(idx)
		if err != nil {
			return
		}
		if item.Name == params["name"] {
			jsonEncode(w, item)
			return
		}
	}
	jsonEncode(w, &Contact{})
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	jsonDecode(r, &contact)
	contacts = append(contacts, contact)
	jsonEncode(w, contact)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		_, err := fmt.Println(idx)
		if err != nil {
			return
		}
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			var contact Contact
			jsonDecode(r, &contact)
			contact.Name = params["name"]
			contacts = append(contacts, contact)
			jsonEncode(w, contact)
			return
		}
	}

}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		_, err := fmt.Println(idx)
		if err != nil {
			return
		}
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			break
		}
	}
	jsonEncode(w, contacts)
}

/*********************
 * Utility functions *
 *********************/

func jsonEncode(w http.ResponseWriter, s any) {
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		return
	}
}

func jsonDecode(r *http.Request, pointer any) {
	err := json.NewDecoder(r.Body).Decode(pointer)
	if err != nil {
		return
	}
}
