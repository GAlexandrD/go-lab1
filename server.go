package main

import (
	"fmt"
	"log"
	"net/http"
)

type TimeData struct {
    Time string json:"time"
}

func main() {
    fmt.Println("Hello, world.")
    log.Fatal(http.ListenAndServe(":8795", nil))
}