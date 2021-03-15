package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Whoa go is neat!</h1>
	<p>Whoa go is neat!</p>
	<p>...and simple!</p>`)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "expert web design by sendex")
}


func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	fmt.Println("server start at http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
