package main

import (
    "fmt"
    "log"
    "net/http"
)

func getupdates(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func postupdates(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/getupdates", getupdates)
    mux.HandleFunc("/postupdates", postupdates)
    log.Fatal(http.ListenAndServe(":8080", mux))
}