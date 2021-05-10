package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// HistoryRequest type for request body
type HistoryRequest struct {
	FileUrl string `json:"file_url"`
}

// Calcultion type for response body
type Calculation struct {
	MaxProfit int `json:"max_profit"`
	BuyHour   int `json:"buy_hour"`
	BuyPrice  int `json:"buy_price"`
	SellHour  int `json:"sell_hour"`
	SellPrice int `json:"sell_price"`
}

func MaxTrade(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: MaxTrade")
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

	dataBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "Error reading response", http.StatusBadRequest)
		return
	}
	bodyString := string(dataBody)
	bodyArray := strings.Split(bodyString, " ")

	arr := make([]int, len(bodyArray))

	for i, v := range bodyArray {
		arr[i], err = strconv.Atoi(v)
		if err != nil {
			http.Error(w, "Error parsing body", http.StatusBadRequest)
			return
		}
	}

	calc, err := CalculateProfit(arr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calc)
}

// CalculateProfit main logic for calculation
func CalculateProfit(prices []int) (*Calculation, error) {
	calc := &Calculation{}
	// initiate temp values
	profit := 0
	buy_idx := 0
	sell_idx := 0
	low_idx := 0

	for i := 0; i < len(prices); i++ {

		if prices[i]-prices[low_idx] > profit {
			buy_idx = low_idx
			sell_idx = i
			profit = prices[i] - prices[buy_idx]
		}
		if prices[i] < prices[low_idx] {
			low_idx = i
		}
	}
	calc.MaxProfit = profit
	calc.BuyHour = buy_idx + 1
	calc.BuyPrice = prices[buy_idx]
	calc.SellHour = sell_idx + 1
	calc.SellPrice = prices[sell_idx]
	return calc, nil
}
