package main

import (
    "net/http"
    "github.com/brettbuddin/dapper"
)

func main() {
    router := dapper.NewRouter()

    router.Get("stream", StreamHandler)

    http.ListenAndServe(":4000", router)
}

func StreamHandler(resp *dapper.Responder, req *dapper.Request) {
    messages := make(chan []byte)

    go func() {
        for {
            messages <- []byte("howdy!")
        }
    }()

    resp.Stream(messages)
}
