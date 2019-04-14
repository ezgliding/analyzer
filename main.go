package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("serving request\n")
		fmt.Fprintln(w, "hello world")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
