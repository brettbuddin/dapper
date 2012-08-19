package main

import (
    "net/http"
    "github.com/brettbuddin/dapper"
)

func main() {
    router := dapper.NewRouter()

    router.Root(HomeHandler)
    router.Get("hello", OtherHandler)

    http.ListenAndServe(":4000", router)
}

func HomeHandler(resp *dapper.Responder, req *dapper.Request) {
    resp.Respond([]byte("home page"), 200)
}

func OtherHandler(resp *dapper.Responder, req *dapper.Request) {
    resp.Respond([]byte("hello"), 200)
}

