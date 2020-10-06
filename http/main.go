package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gabrielnascr/golang-study/http/middleware"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	http.HandleFunc("/", middleware.BasicMiddleware(handler))
	log.Fatal(http.ListenAndServe(":3333", nil))
}
