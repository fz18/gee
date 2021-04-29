package gee

import (
	"net/http"
)

// HandleFunc defines the request handler by gee
type HandleFunc func(c *Context)

type Engine struct {
	router *router
}

// New is a constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// Get method
func (e *Engine) Get(pattern string, handler HandleFunc) {
	e.router.addRouter("GET", pattern, handler)
}

// Post method
func (e *Engine) Post(pattern string, handler HandleFunc) {
	e.router.addRouter("POST", pattern, handler)
}

// Run method
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}
