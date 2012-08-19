package main

import (
    "net/http"
)

func main() {
    router := NewRouter()

    router.Root(HomeHandler)
    router.Get("stream", StreamHandler)

    http.ListenAndServe(":4000", router)
}

func HomeHandler(resp *Responder, req *Request) {
    panic("yoyoooyyoy")
    resp.Respond([]byte("home page"), 200)
}

func StreamHandler(resp *Responder, req *Request) {
    messages := make(chan []byte)

    go func() {
        for {
            messages <- []byte("Test")
        }
    }()

    resp.Stream(messages)
    //resp.Respond([]byte(params["filename"]), 200)
}
