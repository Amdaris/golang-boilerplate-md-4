package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	challenges "backend-boilerplate-golang/challenges"
	_ "backend-boilerplate-golang/docs"
)

// @Tags         default
// @Produce      json
// @Success      200 {string} string "pong"
// @Router       /ping [get]
func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "pong")
}

// @Tags         default
// @Accept       json
// @Produce      json
// @Param        request body []int true "request"
// @Success      200 {boolean} bool "response"
// @Router       /challenge-1 [post]
func challenge1Handler(w http.ResponseWriter, r *http.Request) {
	var array []int
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&array); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := challenges.Challenge1(array)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// @Tags         default
// @Accept       json
// @Produce      json
// @Param        request body []int true "request"
// @Success      200 {array} int "response"
// @Router       /challenge-2 [post]
func challenge2Handler(w http.ResponseWriter, req *http.Request) {
	var array []int
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&array); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := challenges.Challenge2(array)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// @Tags         default
// @Accept       json
// @Produce      json
// @Param        request body []int true "request"
// @Success      200 {integer} int "response"
// @Router       /challenge-3 [post]
func challenge3Handler(w http.ResponseWriter, req *http.Request) {
	var array []int
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&array); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := challenges.Challenge3(array)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// @title           Internship Challenge
// @version         1.0
func main() {
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/challenge-1", challenge1Handler)
	http.HandleFunc("/challenge-2", challenge2Handler)
	http.HandleFunc("/challenge-3", challenge3Handler)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
