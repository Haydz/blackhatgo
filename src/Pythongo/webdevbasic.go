package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
	fmt.Fprintf(w, "<p>Go is fast!</p>")
	fmt.Fprintf(w, "<p>...and simple!</p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong")
	fmt.Fprintf(w, `
	<h6> You can even do ... </h6>
	<hh>miltiple lines...</h5>
	<h4>in one %s</h4>`, "formatted print")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Harrison Kinsley")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
