package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("first r: %v\n", r)
	r.ParseForm()
	fmt.Println("----Http.Request----")
	fmt.Println("----line----")
	fmt.Println(r.Method, r.RequestURI, r.Proto)
	fmt.Println("----header----")
	fmt.Println(r.Header)
	fmt.Println("----body----")
	fmt.Println(r.Body)
	fmt.Println("----end----")
	fmt.Fprintf(w, "hello go server!")


}

func main() {
	http.HandleFunc("/", sayhelloName)
	fmt.Println("Listen in 9090.")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listen and server:", err)
	}
}
