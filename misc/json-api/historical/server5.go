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

type character struct {
	Name  string `json: "name"`
	Race  string `json: "race"`
	Level int    `json: "level"`
}
type charComp func(r *http.Request) ([]byte, error)

func magicHandler(f charComp) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := f(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("There was an error loading : %s", err)
			return
		}
		if _, err := w.Write(b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("There was an error writing: %s", err)
		}
	}
}

func charList(r *http.Request) ([]byte, error) {
	t := []character{
		{Level: 1, Name: "Maloy", Race: "Dwarf"},
		{Level: 10, Name: "Claw", Race: "Mountain Lion"},
		{Level: 19, Name: "Clem", Race: "Elf"},
	}
	return json.Marshal(t)
}

func spellList(r *http.Request) ([]byte, error) {
	t := []spell{
		{Level: 1, Name: "loud", Description: "Double the decibel, but no higher than 11."},
		{Level: 2, Name: "frustrate", Description: "You speak for hours about the liberal agenda"},
	}
	return json.Marshal(t)
}
func main() {

	http.HandleFunc("/api/spells", magicHandler(spellList))
	http.HandleFunc("/api/character", magicHandler(charList))
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
