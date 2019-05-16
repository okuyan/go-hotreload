package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8083"
	http.HandleFunc("/hello", hello)
	log.Printf("Listening on localhost: %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
