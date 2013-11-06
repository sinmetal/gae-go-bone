package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/twitter", hello)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, world</h1>")
}
