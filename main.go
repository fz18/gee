package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>Hello Gee<h1>")
	})
	r.Get("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.Get("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.Get("/asset/*filepath", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})
	r.Post("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	err := r.Run(":9999")
	if err != nil {
		fmt.Println(err)
	}
}
