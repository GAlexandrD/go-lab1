package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    fmt.Println("Hello, world.")
    log.Fatal(http.ListenAndServe(":8795", nil))
}