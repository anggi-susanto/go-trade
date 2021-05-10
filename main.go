package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anggi-susanto/go-trade/usecase"
	"github.com/gorilla/mux"
)

// homePage simple homepage handler
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go trade!")
	fmt.Println("Endpoint Hit: homePage")
}

// handleRequests main simple router
func handleRequests() {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/max-trade", usecase.MaxTrade).Methods("POST")
	router.HandleFunc("/unique-string", usecase.UniqueString).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// main main http server
func main() {
	fmt.Println("Go Trade - Traders Left Hand")
	handleRequests()
}
