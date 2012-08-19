package main

import (
    "net/http"
    "github.com/brettbuddin/dapper"
)

func main() {
    router := dapper.NewRouter()

    router.Get("(?P<filename>.*)/info/refs", Handler)

    http.ListenAndServe(":4000", router)
}

func Handler(resp *dapper.Responder, req *dapper.Request) {
    filename := req.Params["filename"]

    resp.Respond([]byte(filename), 200)
}
