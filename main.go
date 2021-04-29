package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	r := gee.New()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %s\n", req.URL.Path)
	})
	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	err := r.Run(":9999")
	if err != nil {
		fmt.Println(err)
	}
}
