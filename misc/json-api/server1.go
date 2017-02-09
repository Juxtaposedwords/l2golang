package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Maloy is totally into, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", handler)
	PORT := os.Getenv("PORT")
	http.ListenAndServe(":"+PORT, nil)

}
