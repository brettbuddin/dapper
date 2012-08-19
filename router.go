package dapper

import (
    "net/http"
    "regexp"
    "log"
    "os"
)

type Handler func(*Responder, *Request)

type Router struct {
    Routes []*route
    Logger *log.Logger
}

func NewRouter() *Router {
    return &Router{
        Logger: log.New(os.Stdout, "", log.Ldate | log.Ltime),
    }
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    r.dispatch(w, req)
}

func (r *Router) Root(handler Handler) {
    r.Get("", handler)
}

func (r *Router) Get(pattern string, handler Handler) {
    r.addRoute("GET", pattern, handler)
}

func (r *Router) Post(pattern string, handler Handler) {
    r.addRoute("POST", pattern, handler)
}

func (r *Router) Put(pattern string, handler Handler) {
    r.addRoute("PUT", pattern, handler)
}

func (r *Router) Delete(pattern string, handler Handler) {
    r.addRoute("DELETE", pattern, handler)
}

func (r *Router) SetLogger(logger *log.Logger) {
    r.Logger = logger
}

func (r *Router) dispatch(w http.ResponseWriter, req *http.Request) {
    method := req.Method
    url    := req.URL

    for _, route := range r.Routes {
        if route.method == "" || route.method != method {
            continue
        }

        matches := route.matcher.FindAllStringSubmatch(url.Path, -1)

        if len(matches) > 0 {
            params  := route.processMatches(matches)
            handler := func(http.ResponseWriter, *http.Request) {
                route.handler(&Responder{w}, &Request{Request: req, Params: params})
            }

            defer r.rescue(w, req)
            http.HandlerFunc(handler)(w, req)

            return
        }
    }
}

func (r *Router) rescue(w http.ResponseWriter, req *http.Request) {
    if err := recover(); err != nil {
        r.Logger.Printf("Handler error: %s", err)
        w.WriteHeader(500)
    }
}

func (r *Router) addRoute(method, pattern string, handler Handler) {
    exp := r.matcherPattern(pattern)
    r.Routes = append(r.Routes, &route{method: method, matcher: exp, handler: handler})

    r.Logger.Printf("Registered route: /%s/", exp.String())
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
