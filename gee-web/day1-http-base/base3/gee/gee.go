package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandleFunc func(w http.ResponseWriter, r *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandleFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handle HandleFunc) {
	engine.addRoute("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandleFunc) {
	engine.addRoute("POST", pattern, handle)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND：%s\n", req.URL)
	}
}
