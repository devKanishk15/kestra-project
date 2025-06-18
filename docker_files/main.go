package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go Web Server!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

