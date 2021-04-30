package gee

import (
	"log"
	"net/http"
)

// HandleFunc defines the request handler by gee
type HandleFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix string
	middlewares []HandleFunc
	parent *RouterGroup
	engine *Engine
}

// New is a constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Get method
func (group *RouterGroup) Get(pattern string, handler HandleFunc) {
	group.addRouter("GET", pattern, handler)
}

// Post method
func (group *RouterGroup) Post(pattern string, handler HandleFunc) {
	group.addRouter("POST", pattern, handler)
}

// Run method
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

// create router group 
func (group *RouterGroup) Group (prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRouter(method, comp string, handler HandleFunc)  {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handler)		
}
