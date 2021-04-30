package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	r := gee.New()
	r.Get("/index", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>Hello Gee<h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.Get("/index", func(c *gee.Context) {
			c.Html(http.StatusOK, "<h1>Hello Gee<h1>")
		})
		v1.Get("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.Get("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.Get("/asset/*filepath", func(c *gee.Context) {
			c.Json(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
		})
		v2.Post("/login", func(c *gee.Context) {
			c.Json(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	err := r.Run(":9999")
	if err != nil {
		fmt.Println(err)
	}
}
