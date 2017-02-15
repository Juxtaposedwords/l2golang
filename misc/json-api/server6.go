package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func listChars(r *http.Request) ([]byte, error) {
	t := []character{
		{Level: 1, Name: "Maloy", Race: "Dwarf"},
		{Level: 10, Name: "Claw", Race: "Mountain Lion"},
		{Level: 19, Name: "Clem", Race: "Elf"},
	}
	return json.Marshal(t)
}

func getChar(r *http.Request) ([]byte, error) {
	t := []character{
		{Level: 1, Name: "Maloy", Race: "Dwarf"},
		{Level: 10, Name: "Claw", Race: "Mountain Lion"},
		{Level: 19, Name: "Clem", Race: "Elf"},
	}
	char := r.URL.Path[len("/api/character/"):]
	for _, e := range t {
		if e.Name == char {
			return json.Marshal(e)
		}
	}
	return []byte(fmt.Sprintf("The character %s has not been created yet.", char)), nil

}
func listLevelChars(r *http.Request) ([]byte, error) {
	t := []character{
		{Level: 1, Name: "Maloy", Race: "Dwarf"},
		{Level: 10, Name: "Claw", Race: "Mountain Lion"},
		{Level: 19, Name: "Clem", Race: "Elf"},
	}
	s, err := strconv.Atoi(r.URL.Path[len("/api/characters/"):])
	if err != nil {
		return nil, err
	}
	var o []character
	for _, e := range t {
		if e.Level == s {
			o = append(o, e)
		}
	}
	if len(o) != 0 {
		return json.Marshal(o)
	}
	return []byte(fmt.Sprintf("There are no characters at level %d ...yet", s)), nil

}
func listSpells(r *http.Request) ([]byte, error) {
	t := []spell{
		{Level: 1, Name: "loud", Description: "Double the decibel, but no higher than 11."},
		{Level: 1, Name: "distract", Description: "How do spell levels map to character levels again?"},
		{Level: 2, Name: "frustrate", Description: "You speak for hours about the liberal agenda"},
	}
	return json.Marshal(t)
}

func listLevelSpells(r *http.Request) ([]byte, error) {
	t := []spell{
		{Level: 1, Name: "loud", Description: "Double the decibel, but no higher than 11."},
		{Level: 1, Name: "distract", Description: "How do spell levels map to character levels again?"},
		{Level: 2, Name: "frustrate", Description: "You speak for hours about the liberal agenda"},
	}
	s, err := strconv.Atoi(r.URL.Path[len("/api/spells/"):])
	if err != nil {
		return nil, err
	}

	var o []spell

	for _, e := range t {
		if e.Level == s {
			o = append(o, e)
		}
	}
	if len(o) != 0 {
		return json.Marshal(o)
	}
	return []byte(fmt.Sprintf("There are no level %d spells...Some say the magic has gone away.", s)), nil
}

func getSpell(r *http.Request) ([]byte, error) {
	t := []spell{
		{Level: 1, Name: "loud", Description: "Double the decibel, but no higher than 11."},
		{Level: 1, Name: "distract", Description: "How do spell levels map to character levels again?"},
		{Level: 2, Name: "frustrate", Description: "You speak for hours about the liberal agenda"},
	}
	s := r.URL.Path[len("/api/spell/"):]
	for _, e := range t {
		if e.Name == s {
			return json.Marshal(e)
		}
	}
	return []byte(fmt.Sprintf("There is no such spell as %s!", s)), nil

}

func main() {

	http.HandleFunc("/api/spells", magicHandler(listSpells))
	http.HandleFunc("/api/spells/", magicHandler(listLevelSpells))
	http.HandleFunc("/api/spell/", magicHandler(getSpell))
	http.HandleFunc("/api/characters", magicHandler(listChars))
	http.HandleFunc("/api/characters/", magicHandler(listLevelChars))
	http.HandleFunc("/api/character/", magicHandler(getChar))
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
