package main

import (
    "fmt"
    "net/http"
    )

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello")
}

func main() {
    fmt.Println("Hello")
    http.HandleFunc("/home", Home)
    http.ListenAndServe(":8081", nil)
}

