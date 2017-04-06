package myData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type character struct {
	ID    int    `json: "int"`
	Name  string `json: "name"`
	Race  string `json: "race"`
	Level int    `json: "level"`
}
type spell struct {
	ID          int    `json: "id"`
	Level       int    `json: "level"`
	Name        string `json: "name"`
	Description string `json: "description"`
}

func loadJSON(title string) ([]byte, error) {
	fn := "../resources/" + title + ".json"
	return ioutil.ReadFile(fn)
}

func saveJSON(title string, b []byte) error {
	fn := "../resources/" + title + ".json"
	return ioutil.WriteFile(fn, b, 0600)

}

func ListSpells() ([]spell, error) {
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

func ListCharacters() ([]character, error) {
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
func GetSpell(id int) (spell, error) {
	l, err := ListSpells()
	if err != nil {
		return spell{}, err
	}
	for _, e := range l {
		if id == e.ID {
			return e, nil
		}
	}
	return spell{}, fmt.Errorf("Entity not found")

}
func GetCharacter(id int) (character, error) {
	l, err := ListCharacters()
	if err != nil {
		return character{}, err
	}
	for _, e := range l {
		if id == e.ID {
			return e, nil
		}
	}
	return character{}, fmt.Errorf("Entity not found")
}
func PutSpell(input spell) error {
	l, err := ListSpells()
	if err != nil {
		return err
	}
	added := false
	for i, e := range l {
		if input.ID == e.ID {
			l[i] = input
			added = true
		}
	}
	if !added {
		id, err := newUID("spells")
		if err != nil {
			return err
		}
		input.ID = id
		l = append(l, input)
	}
	b, err := json.Marshal(l)
	if err != nil {
		return nil
	}
	err = saveJSON("spells", b)
	if err != nil {
		return err
	}
	return nil
}
func PutCharacter(input character) error {
	l, err := ListCharacters()
	if err != nil {
		return err
	}
	added := false
	for i, e := range l {
		if input.ID == e.ID {
			l[i] = input
			added = true
		}
	}
	id, err := newUID("characters")
	if err != nil {
		return err
	}
	if !added {
		input.ID = id
		l = append(l, input)
	}
	b, err := json.Marshal(l)
	if err != nil {
		return nil
	}
	err = saveJSON("characters", b)
	if err != nil {
		return err
	}
	return nil
}

func newUID(s string) (int, error) {
	// all the stuff to load the map of items
	b, err := loadJSON("meta")
	if err != nil {
		return 0, err
	}
	idMap := map[string]int{}
	if err := json.Unmarshal(b, &idMap); err != nil {
		return 0, err
	}
	_, ok := idMap[s]
	if !ok {
		return 0, fmt.Errorf("there is no entry for %s", s)
	}
	// This is literally the two lines the function does, it gets an int and sets it
	idMap[s]++
	output := idMap[s]

	// all stuff to just save it back
	b, err = json.Marshal(idMap)
	if err != nil {
		return 0, err
	}
	err = saveJSON("meta", b)
	if err != nil {
		return 0, err
	}
	return output, nil
}
