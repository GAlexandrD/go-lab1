package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type TimeData struct {
    Time string json:"time"
}

func getTime(w http.ResponseWriter, _ *http.Request) {
	t := TimeData{Time: time.Now().Format(time.RFC3339)}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
  }

func main() {
    fmt.Println("Hello, world.")
	http.HandleFunc("/time", getTime);
    log.Fatal(http.ListenAndServe(":8795", nil))
}