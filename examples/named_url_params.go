package main

import (
    "net/http"
)

func main() {
    router := NewRouter()

    router.Get("(?P<filename>.*)/info/refs", Handler)

    http.ListenAndServe(":4000", router)
}

func Handler(resp *Responder, req *Request) {
    filename := req.Params["filename"]

    resp.Respond([]byte(filename), 200)
}
