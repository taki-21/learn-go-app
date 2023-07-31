package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome glang server.")
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello guest!")
}

func main() {
	port := "8080"
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/greeting", handleGreet)
	log.Printf("Server listening on http://localhost:%s/", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}
