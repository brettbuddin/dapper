package dapper

import (
    "net/http"
    "regexp"
)

type Router struct {
    routes []*route
}

func NewRouter() *Router {
    return new(Router)
}

func (r *Router) dispatch(w http.ResponseWriter, req *http.Request) {
    method := req.Method
    url    := req.URL

    for _, route := range r.routes {
        if route.method == "" || route.method != method {
            continue
        }

        matches := route.matcher.FindAllStringSubmatch(url.Path, -1)

        if len(matches) > 0 {
            params  := route.processMatches(matches)
            handler := func(http.ResponseWriter, *http.Request) {
                route.handler(&Responder{w}, &Request{Request: req, Params: params})
            }

            http.HandlerFunc(handler)(w, req)
            recover()
        }
    }
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    r.dispatch(w, req)
}

func (r *Router) Root(handler func(*Responder, *Request)) {
    r.Get("", handler)
}

func (r *Router) Get(pattern string, handler func(*Responder, *Request)) {
    exp := r.matcherPattern(pattern)
    r.routes = append(r.routes, &route{method: "GET", matcher: exp, handler: handler})
}

func (r *Router) Post(pattern string, handler func(*Responder, *Request)) {
    exp := r.matcherPattern(pattern)
    r.routes = append(r.routes, &route{method: "POST", matcher: exp, handler: handler})
}

func (r *Router) Put(pattern string, handler func(*Responder, *Request)) {
    exp := r.matcherPattern(pattern)
    r.routes = append(r.routes, &route{method: "PUT", matcher: exp, handler: handler})
}

func (r *Router) Delete(pattern string, handler func(*Responder, *Request)) {
    exp := r.matcherPattern(pattern)
    r.routes = append(r.routes, &route{method: "DELETE", matcher: exp, handler: handler})
}

func (r *Router) matcherPattern(pattern string) *regexp.Regexp {
    exp, _ := regexp.Compile("^/(?:" + pattern + ")$")
    return exp
}

type route struct {
    method  string
    matcher *regexp.Regexp
    handler func(*Responder, *Request)
}

func (r *route) processMatches(matches [][]string) map[string]string {
    names := r.matcher.SubexpNames()
    out   := make(map[string]string, 0)

    for i, value := range names {
        if value != "" {
            out[value] = matches[0][i] 
        }
    }

    return out
}
