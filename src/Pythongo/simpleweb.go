package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!", r)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABOUT PAGE")
}

//setup main fun
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
