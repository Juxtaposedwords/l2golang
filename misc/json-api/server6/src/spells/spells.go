package spells

import (
	"encoding/json"
	"fmt"
	"myHTTP"
	"myJSON"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type spell struct {
	Level       int    `json: "level"`
	Name        string `json: "name"`
	Description string `json: "description"`
}
type handler func(*http.Request) ([]byte, error)

const (
	listSpellPattern      = `^/api/spells$`
	addSpellPattern       = `^/api/spells/add$`
	listSpellLevelPattern = `^/api/spells/\d+$`
)

var dispatch = map[string]handler{
	listSpellPattern:      ListSpells,
	addSpellPattern:       AddSpell,
	listSpellLevelPattern: ListLevelSpells,
}

func Dispatcher(r *http.Request) ([]byte, error) {
	p := r.URL.Path
	for k, v := range dispatch {
		ok, err := regexp.MatchString(k, p)
		if err != nil {
			return nil, err
		}
		if ok {
			return v(r)
		}
	}
	return nil, myHTTP.NotFoundErr
}

func LoadSpells() ([]spell, error) {
	b, err := myJSON.LoadJSON("spells")
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
	err = myJSON.SaveJSON("spells", b)
	if err != nil {
		return err
	}
	return nil
}

func ListSpells(r *http.Request) ([]byte, error) {
	return myJSON.LoadJSON("spells")
}

func ListLevelSpells(r *http.Request) ([]byte, error) {
	t, err := LoadSpells()
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

func GetSpell(r *http.Request) ([]byte, error) {
	t, err := LoadSpells()
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

func AddSpell(r *http.Request) ([]byte, error) {
	_, err := LoadSpells()
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
