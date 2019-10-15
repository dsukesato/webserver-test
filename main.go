package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}


func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello/", handler)

	http.ListenAndServe(":8080", mux)
}