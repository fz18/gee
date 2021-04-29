package main

import (
	"example/gee"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	engine := new(gee.Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
