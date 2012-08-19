package main

import (
    "net/http"
)

func main() {
    router := NewRouter()

    router.Root(HomeHandler)
    router.Get("hello", OtherHandler)

    http.ListenAndServe(":4000", router)
}

func HomeHandler(resp *Responder, req *Request) {
    resp.Respond([]byte("home page"), 200)
}

func OtherHandler(resp *Responder, req *Request) {
    resp.Respond([]byte("hello"), 200)
}

