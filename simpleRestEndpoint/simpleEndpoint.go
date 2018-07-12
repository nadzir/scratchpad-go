package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func simpleEndpointGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Simple endpoint")
}

func simpleEndpointPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Simple endpoint")
}

func simpleEndpointPUT(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Simple endpoint")
}

func simpleEndpointDELETE(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Simple endpoint")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/simpleEndpoint", simpleEndpointGET).Methods("GET")
	r.HandleFunc("/simpleEndpoint", simpleEndpointPOST).Methods("POST")
	r.HandleFunc("/simpleEndpoint", simpleEndpointPUT).Methods("PUT")
	r.HandleFunc("/simpleEndpoint", simpleEndpointDELETE).Methods("DELETE")

	port := ":3100"
	fmt.Println("Listening on port ", port)
	err := http.ListenAndServe(":3100", r)
	if err != nil {
		log.Fatal(err)
	}
}
