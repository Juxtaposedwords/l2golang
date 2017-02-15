package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type spell struct {
	Level       int    `json: "level"`
	Name        string `json: "name"`
	Description string `json: "description"`
}

func spellHandler(w http.ResponseWriter, r *http.Request) {
	b, err := spellList(r)
	if err != nil {
		log.Printf("There was an error loading Spells: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		log.Printf("There was an error writing: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func spellList(r *http.Request) ([]byte, error) {
	t := []spell{
		{Level: 1, Name: "loud", Description: "Double the decibel, but no higher than 11."},
		{Level: 2, Name: "frustrate", Description: "You speak for hours about the liberal agenda"},
	}
	return json.Marshal(t)
}
func main() {

	http.HandleFunc("/api/spells", spellHandler)
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
