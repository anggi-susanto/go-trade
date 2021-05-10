package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

// StringRequest type for request body
type StringRequest struct {
	FileUrl string `json:"file_url"`
}

type UniqueStringResponse struct {
	FirstOccurence    string `json:"first_occurence"`
	LexicoGraphically string `json:"smallest_lexicographical_order"`
}

func UniqueString(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UniqueString")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	request := &StringRequest{}

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

	response := &UniqueStringResponse{}

	response.FirstOccurence = FirstOccurence(bodyString)
	response.LexicoGraphically = LexicoGraphically(bodyString)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func FirstOccurence(str string) string {
	result := ""
	occurence := make([]bool, 256)

	for i := 0; i < len(str); i++ {
		char := str[i]
		if !occurence[char] {
			result += string(char)
			occurence[char] = true
		}
	}
	return result
}

func LexicoGraphically(s string) string {
	r := []rune(s)
	r = RemoveDuplicate(r)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func RemoveDuplicate(sliceData []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range sliceData {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
