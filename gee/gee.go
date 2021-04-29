package gee

import (
	"fmt"
	"net/http"
)

// HandleFunc defines the request handler by gee
type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

// New is a constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// AddRouter method
func (e *Engine) addRouter(method, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// Get method
func (e *Engine) Get(pattern string, handler HandleFunc) {
	e.addRouter("GET", pattern, handler)
}

// Post method
func (e *Engine) Post(pattern string, handler HandleFunc) {
	e.addRouter("Post", pattern, handler)
}

// Run method
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get router
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 not found: %s\n", r.URL)
	}

}
