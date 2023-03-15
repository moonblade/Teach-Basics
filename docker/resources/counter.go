package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Counter struct {
	Value int                `bson:"count"`
}

var count Counter // initialize counter to zero

func initializeCount() error {
  count = Counter{Value: 0}
	return nil
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	// Return the current count value
	json.NewEncoder(w).Encode(count)
}

func increaseHandler(w http.ResponseWriter, r *http.Request) {
	// Increase the count value and return the updated count
	count.Value++
	json.NewEncoder(w).Encode(count)
}

func decreaseHandler(w http.ResponseWriter, r *http.Request) {
	// Decrease the count value and return the updated count
	count.Value--
	json.NewEncoder(w).Encode(count)
}

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		next(w, r)
	}
}

func main() {
	err := initializeCount()
	if err != nil {
		fmt.Println("Failed to initialize count:", err)
		os.Exit(1)
	}

	http.HandleFunc("/count", enableCors(countHandler))
	http.HandleFunc("/increase", enableCors(increaseHandler))
	http.HandleFunc("/decrease", enableCors(decreaseHandler))

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
