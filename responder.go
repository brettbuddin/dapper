package dapper

import (
    "net/http"
)

type Responder struct {
    http.ResponseWriter
}

func (r *Responder) SetHeader(key, val string) {
    r.Header().Set(key, val)
}

func (r *Responder) WriteString(content string) {
    r.Write([]byte(content))
}

func (r *Responder) Respond(buf []byte, code int) {
    r.WriteHeader(code)
    r.Write(buf)
}

func (r *Responder) Stream(source chan []byte) {
    for in := range source {
        r.Write(append(in, '\n'))
    }
}
