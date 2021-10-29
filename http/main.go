package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	data := "hello world"
	fmt.Fprintf(w, data)
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic("server error")
	}
}
