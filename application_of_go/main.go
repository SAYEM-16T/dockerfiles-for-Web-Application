package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome!")
    })

    http.HandleFunc("/how are you", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "I am good, how about you?")
    })

    fmt.Println("Server starting on http://0.0.0.0:8080")
    http.ListenAndServe("0.0.0.0:8080", nil)
}