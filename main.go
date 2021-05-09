package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// HistoryRequest type for request body
type HistoryRequest struct {
	FileUrl string `json:"file_url"`
}

// Calcultion type for response body
type Calculation struct {
	MaxProfit int `json:"max_profit"`
}

// homePage simple homepage handler
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go trade!")
	fmt.Println("Endpoint Hit: homePage")
}

// historyCalculate history calculation controller
func historyCalculate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: historyCalculate")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	request := &HistoryRequest{}

	err = json.Unmarshal(body, request)
	if err != nil {
		log.Printf("Error casting body values: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	if request.FileUrl == "" {
		http.Error(w, "file_url must not empty", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(request.FileUrl)

	if err != nil {
		http.Error(w, "error getting file url", http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	calc, err := handleCalculation(resp)

	json.NewEncoder(w).Encode(calc)
}

// handleCalculation file read handler
func handleCalculation(responseFile *http.Response) (*Calculation, error) {
	calc := &Calculation{}
	body, err := ioutil.ReadAll(responseFile.Body)

	if err != nil {
		fmt.Println("Error reading response", err.Error())
		return calc, err
	}

	bodyString := string(body)

	diff, err := processBody(bodyString)
	calc.MaxProfit = diff
	return calc, err
}

// processBody main logic for calculation
func processBody(bodyString string) (int, error) {
	arr := strings.Split(bodyString, " ")
	firstArr, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, err
	}
	// initiate temp values
	min := firstArr
	max := firstArr
	diff := 0

	for _, s := range arr {
		val, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}

		// whenever get lower price it will reset all values
		if min > val {
			min = val
			max = val
			diff = 0
		}

		// whenever got max value change max value
		if max < val {
			max = val
		}

		// assining new diff
		newDiff := max - min
		if newDiff > diff {
			diff = newDiff
		}
	}
	return diff, nil
}

// handleRequests main simple router
func handleRequests() {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/calculate", historyCalculate).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// main main http server
func main() {
	fmt.Println("Go Trade - Traders Left Hand")
	handleRequests()
}
