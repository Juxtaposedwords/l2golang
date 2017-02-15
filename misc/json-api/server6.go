package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

// func magicJSON(title string) ([]interface{}, error) {
// 	fn := "resources/" + title + ".json"
// 	b, err := ioutil.ReadFile(fn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	m := map[string]interface{}{
// 		{"characters": character{},
// 		{"spells": spell{}},
// 	}
// 	c := m[title]
// 	o := []c{}
// 	if err := json.Unmarshal(b, &o); err != nil {
// 		return nil, err
// 	}

// 	return b, nil
// }

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

func loadJSON(title string) ([]byte, error) {
	fn := "resources/" + title + ".json"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func saveJSON(title string, b []byte) error {
	fn := "resources/" + title + ".json"
	err := ioutil.WriteFile(fn, b, 0600)
	if err != nil {
		return err
	}
	return nil

}

func loadSpells() ([]spell, error) {
	b, err := loadJSON("spells")
	if err != nil {
		return nil, err
	}
	var s []spell
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func saveSpells(s []spell) error {
	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	err = saveJSON("spells", b)
	if err != nil {
		return err
	}
	return nil
}

func loadChars() ([]character, error) {
	b, err := loadJSON("characters")
	if err != nil {
		return nil, err
	}
	var s []character
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func saveChars(c []character) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = saveJSON("chracters", b)
	if err != nil {
		return err
	}
	return nil
}

func listChars(r *http.Request) ([]byte, error) {
	return loadJSON("characters")
}

func getChar(r *http.Request) ([]byte, error) {
	x, err := loadChars()
	if err != nil {
		return nil, err
	}
	char := r.URL.Path[len("/api/character/"):]
	for _, e := range x {
		if e.Name == char {
			return json.Marshal(e)
		}
	}
	return []byte(fmt.Sprintf("The character %s has not been created yet.", char)), nil

}
func listLevelChars(r *http.Request) ([]byte, error) {
	t, err := loadChars()
	if err != nil {
		return nil, err
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
	return loadJSON("spells")
}

func listLevelSpells(r *http.Request) ([]byte, error) {
	t, err := loadSpells()
	if err != nil {
		return nil, err
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
	t, err := loadSpells()
	if err != nil {
		return nil, err
	}
	s := r.URL.Path[len("/api/spell/"):]
	for _, e := range t {
		if e.Name == s {
			return json.Marshal(e)
		}
	}
	return []byte(fmt.Sprintf("There is no such spell as %s!", s)), nil
}

func addSpell(r *http.Request) ([]byte, error) {
	_, err := loadSpells()
	if err != nil {
		return nil, err
	}
	if r.Method != "POST" {
		return nil, fmt.Errorf("Attempted to GET, not PUT")
	}
	r.ParseForm()
	fmt.Printf("%v", r.Form["level"])

	_, err = strconv.Atoi(strings.TrimSpace(strings.Join(r.Form["level"], "")))
	if err != nil {
		return nil, err
	}
	// n := spell{
	// 	Level:       l,
	// 	Name:        r.Form["name"][1],
	// 	Description: r.Form["description"][1],
	// }
	// t = append(t, n)
	// err = saveSpells(t)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func main() {

	http.HandleFunc("/api/spells", magicHandler(listSpells))
	http.HandleFunc("/api/spells/", magicHandler(listLevelSpells))
	http.HandleFunc("/api/spell/", magicHandler(getSpell))
	http.HandleFunc("/api/spell/add", magicHandler(addSpell))
	http.HandleFunc("/api/characters", magicHandler(listChars))
	http.HandleFunc("/api/characters/", magicHandler(listChars))
	http.HandleFunc("/api/character/", magicHandler(getChar))
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
