package gee

import "fmt"

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

func (r *router) handle(c *Context) {
	// get router
	key := c.Req.Method + "-" + c.Req.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "404 not found: %s\n", c.Req.URL)
	}
}

// AddRouter method
func (r *router) addRouter(method, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}
