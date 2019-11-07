package main

import (
    "fmt"
    //"log"
    "net/http"
    "encoding/json"
    "serene-light-api/src/common"
    "os"
    "os/signal"
    "syscall"
    "serene-light-api/src/srv/game"
)

var quit bool
var world game.World

func getupdates(w http.ResponseWriter, r *http.Request) {

    m := message.Message{"Alice", "Hello", 1294706395881547000}
    b, _ := json.Marshal(m)

    fmt.Fprintf(w, string(b))
}

func postupdates(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func gameloop(){
    for true{

        if(quit){
            return
        }
    }
}

func main() {

    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    mux := http.NewServeMux()
    mux.HandleFunc("/getupdates", getupdates)
    mux.HandleFunc("/postupdates", postupdates)

    go func() {
        <-c
        quit = true
        os.Exit(1)
    }()

    go http.ListenAndServe(":8080", mux)

    world.Generate()
    gameloop()


}