package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from alpha service")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting alpha service on port 5000")
    http.ListenAndServe(":5000", nil)
}
