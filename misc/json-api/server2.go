package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Maloy Is totally into: %s", r.URL.Path[1:])
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	b, err := someApiFunction(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		log.Printf("error writing resposne: %s", err)
	}

}

func someApiFunction(r *http.Request) ([]byte, error) {
	x := map[string]string{
		"Claw":           "Maloy is crazy for Claw.",
		"Shadowy Mentor": "Maloy is frustrating the mentor by taking too long on his homework.",
		"Golang":         "Maloy is struggle-busing Go.",
	}
	if val, ok := x[r.URL.Path[len("/api/"):]]; ok {
		return []byte(val), nil
	} else {
		return nil, fmt.Errorf("Entry not found")
	}

}
func main() {

	http.HandleFunc("/api/", apiHandler)
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}
