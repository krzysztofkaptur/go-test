package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handleHome)

	http.ListenAndServe(":5000", router)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	type response struct{
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(response{Message: "this is sparta"})
}