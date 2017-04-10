package myData

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"myThings"
	"os"
)

var (
	ErrNotFound = errors.New("Not found.")
	InvalidType = errors.New("Invalid sstruct type.")
)

// reads unmarshals a JSON-tagged struct from an io.Reader
func read(t interface{}, r io.Reader) error {
	b := &bytes.Buffer{}
	if _, err := b.ReadFrom(r); err != nil {
		return err
	}
	return json.Unmarshal(b.Bytes(), t)
}

// write marshals a JSON-tagged struct into a byte streamd and writes it an
//   io.Writer
func write(t interface{}, w io.Writer) error {
	p, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = w.Write(p)
	return err
}
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// read any of hte struct resources
func GetResource(t interface{}, u int) (interface{}, error) {
	fn := "../resources/"
	var x interface{}
	switch typeof(t) {
	case "Character":
		fn += fmt.Sprintf("characters/%d.json", u)
		x = myThings.Character{}
	case "Spell":
		fn += fmt.Sprintf("spells/%d.json", u)
		x = myThings.Spell{}
	default:
		return nil, InvalidType
	}
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	if err = read(x, f); err != nil {
		return nil, err
	}
	return x, nil
}

func loadJSON(title string) ([]byte, error) {
	fn := "../resources/" + title + ".json"
	return ioutil.ReadFile(fn)
}

func saveJSON(title string, b []byte) error {
	fn := "../resources/" + title + ".json"
	return ioutil.WriteFile(fn, b, 0600)

}

func ListSpells() ([]myThings.Spell, error) {
	b, err := loadJSON("spells")
	if err != nil {
		return nil, err
	}
	var s []myThings.Spell
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func ListCharacters() ([]myThings.Character, error) {
	b, err := loadJSON("characters")
	if err != nil {
		return nil, err
	}
	var s []myThings.Character
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}
func GetSpell(id int) (myThings.Spell, error) {
	l, err := ListSpells()
	if err != nil {
		return myThings.Spell{}, err
	}
	for _, e := range l {
		if id == e.ID {
			return e, nil
		}
	}
	return myThings.Spell{}, fmt.Errorf("Entity not found")

}
func GetCharacter(id int) (myThings.Character, error) {
	l, err := ListCharacters()
	if err != nil {
		return myThings.Character{}, err
	}
	for _, e := range l {
		if id == e.ID {
			return e, nil
		}
	}
	return myThings.Character{}, fmt.Errorf("Entity not found")
}
func PutSpell(input myThings.Spell) error {
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

func PutCharacter(input myThings.Character) error {
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
