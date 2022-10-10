package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type property struct {
	gorm.Model
	ID        string `json:"id"`
	Address   string `json:"address"`
	OwnerName string `json:"ownername"`
	Price     int    `json:"price"`
}

var db *gorm.DB
var err error

const DNS = "root:1234@tcp(localhost:3306)/properties"

func getprops(w http.ResponseWriter, r *http.Request) {

}

func addprop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prop property
	json.NewDecoder(r.Body).Decode(&prop)
	db.Create(&property)
	json.NewEncoder(w).Encode(property)
}

func updateprop(w http.ResponseWriter, r *http.Request) {

}

func deleteprop(w http.ResponseWriter, r *http.Request) {

}

func intilializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/props", getprops).Methods("GET")
	r.HandleFunc("/addprop", addprop).Methods("POST")
	r.HandleFunc("/updateprop", updateprop).Methods("PUT")
	r.HandleFunc("/deleteprop", deleteprop).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	intilializeRouter()
	db, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database.")
	}
	db.AutoMigrate(&property)
}
