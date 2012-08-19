package main

import (
    "net/http"
)

func main() {
    router := NewRouter()

    router.Get("stream", OtherHandler)

    http.ListenAndServe(":4000", router)
}

func StreamHandler(resp *Responder, req *Request) {
    messages := make(chan []byte)

    go func() {
        for {
            messages <- []byte("howdy!")
        }
    }()

    resp.Stream(messages)
}
