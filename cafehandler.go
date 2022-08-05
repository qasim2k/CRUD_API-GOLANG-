package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Createcafe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var emp Cafe
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}
func Getcafe(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var employees []Cafe
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)

}

func GetcafebyID(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var employee Cafe
	Database.Find(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(employee)

}

func Updatecafe(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var employee Cafe
	Database.Find(&employee, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)
	json.NewEncoder(w).Encode(employee)

}

func DeleteCafe(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
	var employee Cafe
	Database.Delete(&employee, mux.Vars(r)["eid"])

	json.NewEncoder(w).Encode("Record deleted successfully....")

}
