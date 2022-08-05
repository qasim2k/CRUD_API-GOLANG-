package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/cafe_add", Createcafe).Methods("POST")
	r.HandleFunc("/getcafe", Getcafe).Methods("GET", "OPTIONS")
	r.HandleFunc("/cafebyid/{eid}", GetcafebyID).Methods("GET")
	r.HandleFunc("/deletecafe/{eid}", DeleteCafe).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/updatecafe/{eid}", Updatecafe).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))
}
