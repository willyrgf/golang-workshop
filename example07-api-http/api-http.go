package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "world"
	found := r.URL.Query().Get("name")
	if len(found) > 0 {
		msg = found
	}

	fmt.Fprintf(w, "Hello, %s!\nURL path: %s\n", msg, r.URL.Path)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
