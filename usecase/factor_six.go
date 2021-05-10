package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FactorRequest struct {
	Number int `json:"number"`
}

type FactorResponse struct {
	MatchedFactorCount int `json:"mathched_factor_count"`
}

func FactorSixCount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: FactorSixCount")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	request := &FactorRequest{}

	err = json.Unmarshal(body, request)
	if err != nil {
		log.Printf("Error casting body values: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	response := &FactorResponse{}

	response.MatchedFactorCount = FindFactorSixCount(request.Number)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func factors(n int) []int {
	res := []int{}
	for t := n; t > 0; t-- {
		if (n/t)*t == n {
			res = append(res, t)
		}
	}
	return res
}

func FindFactorSixCount(num int) int {
	result := 0
	matchFactor := 6
	for i := 1; i <= num; i++ {
		factorList := factors(i)
		if len(factorList) == matchFactor {
			result++
		}
	}
	return result
}
