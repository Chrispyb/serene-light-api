package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "serene-light-api/src/common"
)


func getupdates(w http.ResponseWriter, r *http.Request) {

    m := message.Message{"Alice", "Hello", 1294706395881547000}
    b, _ := json.Marshal(m)

    fmt.Fprintf(w, string(b))
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